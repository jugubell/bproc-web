/*
 * File: main.go
 * Project: bproc-web
 * Last modified: 2025-11-24 22:20
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

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
	}))

	router.GET("api", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "velkommen",
		})
	})

	routerGet("api", "help", router)
	routerGet("api", "instruction-set", router)
	routerGet("api", "version", router)
	routerGet("api", "is", router)

	router.POST("/api/verify", func(c *gin.Context) {
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

		var args = []string{
			"-jar",
			"libs/bproc-cli-v1_0.jar",
			"-s",
			"./tmp/prog.bpasm",
		}

		cmd := exec.Command("java", args...)
		out, err := cmd.CombinedOutput()

		var msgType string
		var msgContent string

		if err != nil {
			msgType = "error"
			msgContent = fmt.Sprintf("%s", err)
		} else {
			msgType = "info"
			msgContent = fmt.Sprintf("%s", out)
		}
		c.JSON(200, gin.H{
			"message": msgContent,
			"type":    msgType,
		})

	})

	err := router.Run("localhost:8998")
	if err != nil {
		return
	} // listens on 0.0.0.0:8080 by default
}

func routerGet(prefix string, route string, router *gin.Engine) {
	router.GET(fmt.Sprintf("%s/%s", prefix, route), func(c *gin.Context) {
		var args = []string{
			"-jar",
			"libs/bproc-cli-v1_0.jar",
			fmt.Sprintf("--%s", route),
		}
		cmd := exec.Command("java", args...)
		out, err := cmd.CombinedOutput()

		var msgType string
		var msgContent string

		if err != nil {
			msgType = "error"
			msgContent = fmt.Sprintf("%s", err)
		} else {
			msgType = "info"
			msgContent = fmt.Sprintf("%s", out)
		}
		c.JSON(200, gin.H{
			"message": msgContent,
			"type":    msgType,
		})
	})
}

type ProgramPayload struct {
	Program string `json:"program" binding:"required"`
}
