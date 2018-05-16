package main

import (
	"github.com/gin-gonic/gin"
	"github.com/fvbock/endless"
	"test"
)

func main() {
	router := gin.Default()
	testGroup := router.Group("/api/testGroup/")
	{
		testGroup.GET("/ping", test.PingPongTest)
		testGroup.GET("/login", test.GetLoginTest)
		testGroup.POST("/login", test.PostLoginTest)
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

