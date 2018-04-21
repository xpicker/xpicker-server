package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
)

func main() {
	router := gin.Default()
	test := router.Group("/api/test/")
	{
		test.GET("/ping", PingPongTest)
	}
	endless.ListenAndServe(":10086", router)
}

func PingPongTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
