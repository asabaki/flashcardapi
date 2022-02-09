package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func New() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/api/callback", func(c *gin.Context) {
		data, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(data))
	})
	router.Run()
}
