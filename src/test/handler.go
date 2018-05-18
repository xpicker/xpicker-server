package test

import (
	"github.com/gin-gonic/gin"
	"app"
	"lib"
)

func PingPongTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetLoginTest(c *gin.Context) {
	/*
	获取登录用户信息
	 */
	user := app.GetCheckUserTest()

	c.JSON(200, gin.H{
		"id":            user.Id,
		"password":      user.Password,
		"username":      user.Username,
		"register_time": user.RegisterTime,
		"last_time":     user.LastTime,
		"email":         user.Email,
		"mobile":        user.Mobile,
		"type":          user.Type,
	})
}

func PostLoginTest(c *gin.Context) {
	/*
	登录用户测试
	 */
	username := c.PostForm("username")
	email := c.PostForm("email")
	password := c.PostForm("password")

	user := app.CheckUser(username, email, password)
	if user.Id != "1" {
		c.JSON(200, gin.H{
			"status":  "Login Failed",
			"message": "User Not Found",
		})
	} else {
		cookie, cookieType := lib.GetLoginCookieHash(user.Username)
		app.RedisSet(cookieType, cookie, "EX", "1800")
		c.SetCookie(cookieType, cookie, 1800, "/", "127.0.0.1", false, true)

		c.JSON(200, gin.H{
			"id":            user.Id,
			"username":      user.Username,
			"register_time": user.RegisterTime,
			"last_time":     user.LastTime,
			"email":         user.Email,
			"mobile":        user.Mobile,
			"type":          user.Type,
		})
	}

}
