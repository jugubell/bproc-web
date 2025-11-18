/*
 * File: main.go
 * Project: bproc-web
 * Last modified: 2025-11-18 23:11
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
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "velkommen",
		})
	})

	routerGet("help", router)
	routerGet("instruction-set", router)
	routerGet("version", router)

	err := router.Run("localhost:8998")
	if err != nil {
		return
	} // listens on 0.0.0.0:8080 by default
}

func routerGet(route string, router *gin.Engine) {
	router.GET(route, func(c *gin.Context) {
		var args = []string{
			"-jar",
			"libs/bproc-cli-v1_0.jar",
			fmt.Sprintf("--%s", route),
		}
		cmd := exec.Command("java", args...)
		out, err := cmd.CombinedOutput()

		var msg string
		var msgContent string

		if err != nil {
			msg = "\"error\""
			msgContent = fmt.Sprintf("\"%s\"", err)
		} else {
			msg = "\"message\""
			msgContent = fmt.Sprintf("\"%s\"", out)
		}
		c.JSON(200, gin.H{
			msg: msgContent,
		})
	})
}
