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
		test.GET("/login", GetLoginTest)
		test.POST("/login", PostLoginTest)
	}
	user := router.Group("/api/user/")
	{
		user.GET("/login", GetLoging)
		user.GET("/register", GetRegister)

		user.POST("/login", PostLoging)
		user.POST("/register", PosRegister)
	}

	endless.ListenAndServe(":10086", router)
}

