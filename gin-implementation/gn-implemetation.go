package main

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
)

type PrintJob struct {
	Username int `json:"username"`
	Userage  int `json:"userage"`
}

func main() {
	router := gin.Default()

	router.GET("/details", func(c *gin.Context) {
		c.JSON(200, gin.H{"operingSystemName": runtime.GOOS})
	})

	router.POST("/print", func(c *gin.Context) {
		var data PrintJob
		c.ShouldBindJSON(&data)
		c.JSON(200, gin.H{"userInfo": data.Username})
	})

	router.PUT("/print/:id", func(c *gin.Context) {
		item := c.Param("id")
		var data PrintJob
		c.ShouldBindJSON(&data)
		fmt.Println(data.Username)
		fmt.Println(data.Userage)
		c.JSON(200, gin.H{"userId": item, "name": data.Username, "age": data.Userage})
	})

	router.Run(":4000")
}
