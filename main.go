/*
 * File: main.go
 * Project: bproc-web
 * Last modified: 2025-11-25 21:42
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
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var apiPrefix string = "api"
var host string = "localhost"
var port string = "8998"
var originhost string = "http://localhost"
var originport string = "5173"
var jarName string = "bproc-cli-v1_0.jar"

type ProgramPayload struct {
	Program string `json:"program" binding:"required"`
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{fmt.Sprintf("%s:%s", originhost, originport)},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	router.GET("/", getRoot)
	router.GET(fmt.Sprintf("/%s", apiPrefix), getApiHelp)
	router.GET(fmt.Sprintf("/%s/help", apiPrefix), func(c *gin.Context) { getInfo(c, "--help") })
	router.GET(fmt.Sprintf("/%s/is", apiPrefix), func(c *gin.Context) { getInfo(c, "--instruction-set") })
	router.GET(fmt.Sprintf("/%s/version", apiPrefix), func(c *gin.Context) { getInfo(c, "--version") })

	router.POST(fmt.Sprintf("/%s/verify", apiPrefix), postVerify)

	// run server
	err := router.Run(fmt.Sprintf("%s:%s", host, port))
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
		fmt.Sprintf("libs/%s", jarName),
	}

	cliArgs = append(cliArgs, args...)

	fmt.Println(cliArgs)

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

func postVerify(c *gin.Context) {
	var prog ProgramPayload

	if err := c.ShouldBindJSON(&prog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"type":    "error",
		})
		return
	}

	err := os.MkdirAll("tmp", 0o755)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error creating tmp directory: %v\n", os.Stderr))
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Error creating tmp directory: %v\n. Error: %s", os.Stderr, err),
			"type":    "error",
		})
		return
	}

	f, err := os.Create("tmp/prog.bpasm")
	if err != nil {
		fmt.Println(fmt.Sprintf("Error creating tmp file: %v\n", os.Stderr))
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Error creating tmp file: %v\n. Error: %s", os.Stderr, err),
			"type":    "error",
		})
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(fmt.Sprintf("Error closing tmp file: %v\n", os.Stderr))
		}
	}(f)

	fcontent := []byte(prog.Program)

	if _, err := f.Write(fcontent); err != nil {
		fmt.Println(fmt.Sprintf("Error creating tmp file: %v\n. Error: %s", os.Stderr, err))
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Error creating tmp file: %v\n. Error: %s", os.Stderr, err),
			"type":    "error",
		})
	}

	msgType, msgContent := prepareExec("-s", fmt.Sprintf("%s", "./tmp/prog.bpasm"))

	c.JSON(http.StatusOK, gin.H{
		"message": msgContent,
		"type":    msgType,
	})
}
