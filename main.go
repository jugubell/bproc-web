/*
 * File: main.go
 * Project: bproc-web
 * Last modified: 2025-11-28 23:50
 *
 * This file: main.go is part of BProC-WEB project.
 *
 * BProC-WEB is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published
 * by the Free Software Foundation, either version 2 of the License,
 * or (at your option) any later version.
 *
 * BProC-WEB is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty
 * of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with BProC-WEB. If not, see <https://www.gnu.org/licenses/>.
 *
 * Copyright (C) 2025 Jugurtha Bellagh
 */

package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"slices"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	AppEnv             string
	ApiPrefix          string
	ApiHost            string
	ApiPort            string
	OriginHost         string
	OriginPort         string
	StaticPath         string
	JarPath            string
	ExamplesPath       string
	AllowedCompileType = []string{"bin", "hex", "hexv3", "vhdl", "vrlg"}
)

type ProgramPayload struct {
	Program string `json:"program" binding:"required"`
	Type    string `json:"type"`
}

func main() {
	//loading env vars
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	AppEnv = os.Getenv("APP_ENV")
	ApiPrefix = os.Getenv("API_PREFIX")
	ApiHost = os.Getenv("API_HOST")
	ApiPort = os.Getenv("API_PORT")
	OriginHost = os.Getenv("ORIGIN_HOST")
	OriginPort = os.Getenv("ORIGIN_PORT")
	StaticPath = os.Getenv("STATIC_PATH")
	JarPath = os.Getenv("JAR_PATH")
	ExamplesPath = os.Getenv("EXAMPLES_PATH")

	var router *gin.Engine

	if AppEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)

		router = gin.New()
		router.Use(gin.Recovery())

		log.Println("Gin in Release mode.")

		if err := router.SetTrustedProxies([]string{"127.0.0.1", "::1", "192.168.0.0/16", "172.16.0.0/12", "10.0.0.0/8"}); err != nil {
			log.Println("Failed to set trusted proxies: ", err)
		}

	} else {
		router = gin.Default()
		log.Println("Gin in Default mode.")
		router.Use(cors.New(cors.Config{
			AllowOrigins: []string{fmt.Sprintf("%s:%s", OriginHost, OriginPort)},
			AllowMethods: []string{"GET", "POST"},
			AllowHeaders: []string{"Origin", "Content-Type"},
		}))
		log.Printf("Allowed origin: %s on port %s\n", OriginHost, OriginPort)
	}

	api := router.Group(ApiPrefix)
	{
		api.GET("/", getApiHelp)
		api.GET("/help", func(c *gin.Context) { getInfo(c, "--help") })
		api.GET("/is", func(c *gin.Context) { getInfo(c, "--instruction-set") })
		api.GET("/version", func(c *gin.Context) { getInfo(c, "--version") })
		api.GET("/example", readProgramExample)

		api.POST("/verify", func(c *gin.Context) { postCompile(c, "verify") })
		api.POST("/compile", func(c *gin.Context) { postCompile(c, "compile") })
	}

	// serving static on prod
	if AppEnv == "prod" {

		router.GET("/", func(c *gin.Context) {
			c.File(filepath.Join(StaticPath, "index.html"))
		})

		router.NoRoute(func(c *gin.Context) {
			requestedFile := c.Request.URL.Path
			filePath := filepath.Join(StaticPath, requestedFile)
			_, err := os.Stat(filePath)

			if err == nil {
				c.File(filePath)
				return
			}

			c.File(filepath.Join(StaticPath, "index.html"))
		})
		log.Println("Running on prod http://" + ApiHost + ":" + ApiPort)
	} else {
		log.Println("Running in dev mode on:")
		log.Println("Backend: http://" + ApiHost + ":" + ApiPort)
		log.Println("Frontend: " + OriginHost + ":" + OriginPort)
	}

	// run server
	err := router.Run(fmt.Sprintf("%s:%s", ApiHost, ApiPort))
	if err != nil {
		return
	}
}

func prepareExec(args ...string) (string, string) {
	out, err := execCli(args...)

	var msgType string
	var msgContent string

	if err != nil {
		msgType = "error"
		msgContent = fmt.Sprintf("%s", err)
	} else {
		msgType = "info"
		msgContent = fmt.Sprintf("%s", out)
	}
	return msgType, msgContent
}

func execCli(args ...string) ([]byte, error) {
	var cliArgs = []string{
		"-jar",
		JarPath,
	}

	cliArgs = append(cliArgs, args...)

	log.Printf("Args length: %s\n", len(cliArgs))
	for _, arg := range cliArgs {
		log.Printf("Arg: %s\n", arg)
	}

	cmd := exec.Command("java", cliArgs...)
	out, err := cmd.CombinedOutput()
	return out, err
}

func getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Velkommen",
		"go-to":   "/api",
	})
}

func getApiHelp(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"version": "1.0.0",
		"help": gin.H{
			"prefix": "api",
			"routes": gin.H{
				"help":            "displays help",
				"version":         "displays version",
				"instruction-set": "displays the supported instruction set",
			},
			"route-aliases": gin.H{
				"instruction-set": []string{"is"},
			},
		},
	})
}

func getInfo(c *gin.Context, info string) {
	msgType, msgContent := prepareExec(fmt.Sprintf("%s", info))
	c.JSON(http.StatusOK, gin.H{
		"message": msgContent,
		"type":    msgType,
	})
}

func postCompile(c *gin.Context, action string) {
	var prog ProgramPayload
	var msgContent string
	var msgType string

	if err := c.ShouldBindJSON(&prog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"type":    "error",
		})
		return
	}

	log.Printf("Program type: %s\n", prog.Type)

	if !slices.Contains(AllowedCompileType, prog.Type) && action != "verify" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("%s type is not allowed", prog.Type),
			"type":    "error",
		})
		return
	}

	err := os.MkdirAll("tmp", 0o755)
	if err != nil {
		log.Println(fmt.Sprintf("Error creating tmp directory: %v\n", os.Stderr))
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Error creating tmp directory: %v\n. Error: %s", os.Stderr, err),
			"type":    "error",
		})
		return
	}

	f, err := os.Create("tmp/prog.bpasm")
	if err != nil {
		log.Println(fmt.Sprintf("Error creating tmp file: %v\n", os.Stderr))
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Error creating tmp file: %v\n. Error: %s", os.Stderr, err),
			"type":    "error",
		})
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(fmt.Sprintf("Error closing tmp file: %v\n", os.Stderr))
		}
	}(f)

	fcontent := []byte(prog.Program)

	if _, err := f.Write(fcontent); err != nil {
		log.Println(fmt.Sprintf("Error creating tmp file: %v\n. Error: %s", os.Stderr, err))
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Error creating tmp file: %v\n. Error: %s", os.Stderr, err),
			"type":    "error",
		})
	}

	switch action {
	case "verify":
		msgType, msgContent = prepareExec("-s", fmt.Sprintf("%s", "./tmp/prog.bpasm"))
		break
	case "compile":
		msgType, msgContent = prepareExec("-g", fmt.Sprintf("%s", "./tmp/prog.bpasm"), fmt.Sprintf("--%s", prog.Type))
		break
	}

	c.JSON(http.StatusOK, gin.H{
		"message": msgContent,
		"type":    msgType,
	})
}

func readProgramExample(c *gin.Context) {
	var msgContent string
	var msgType string
	var filesLen int
	var randIdx int

	files, err := os.ReadDir(ExamplesPath)
	if err != nil {
		log.Fatalf("Error reading the directory: %s\n", err)
	}

	filesLen = len(files)
	randIdx = rand.Intn(filesLen)

	log.Printf("%d file(s) found. Selected index: %d. Selected file name: %s\n", filesLen, randIdx, files[randIdx].Name())

	data, err := os.ReadFile(fmt.Sprintf("%s/%s", ExamplesPath, files[randIdx].Name()))

	if err != nil {
		log.Println(fmt.Sprintf("Error reading example file: %v\n", os.Stderr))
		msgContent = ""
		msgType = "error"
	}

	log.Println("File example read.")
	msgContent = string(data)
	msgType = "info"

	c.JSON(http.StatusOK, gin.H{
		"message": msgContent,
		"type":    msgType,
	})
}
