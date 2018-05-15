package main

import (
	"github.com/gin-gonic/gin"
)

func PingPongTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetLoginTest(c *gin.Context) {
	user := CheckUserTest()

	cookie, CookieType := GetLoginCookieHash(user.Username)
	RedisSet(CookieType, cookie, "EX", "1800")
	c.SetCookie(CookieType, cookie, 1800,"/", "", true, false)

	c.JSON(200, gin.H{
		"id": user.Id,
		"password": user.Password,
		"username": user.Username,
		"register_time": user.RegisterTime,
		"last_time": user.LastTime,
		"email": user.Email,
		"mobile": user.Mobile,
		"type": user.Type,
	})
}

func PostLoginTest(c *gin.Context) {
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := CheckUser(username, email, password)

	cookie, CookieType := GetLoginCookieHash(user.Username)
	RedisSet(CookieType, cookie, "EX", "1800")
	c.SetCookie(CookieType, cookie, 1800,"/", "", true, false)

	c.JSON(200, gin.H{
		"id": user.Id,
		"username": user.Username,
		"register_time": user.RegisterTime,
		"last_time": user.LastTime,
		"email": user.Email,
		"mobile": user.Mobile,
		"type": user.Type,
	})
}
func GetLoging(c *gin.Context) {

}
func PostLoging(c *gin.Context) {

}
func GetRegister(c *gin.Context) {

}
func PosRegister(c *gin.Context) {

}