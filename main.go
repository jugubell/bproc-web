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
