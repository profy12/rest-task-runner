package main

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/run/:task", runTask)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func runTask(c *gin.Context) {
	task := c.Params.ByName("task")
	log.Printf("Ask to run task %v", task)
	cmd := exec.Command("task", task)
	out, err := cmd.Output()
	if err != nil {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Unable to run task: %v", err),
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Return " + string(out),
		})
	}
}
