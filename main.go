package main

import (
	"log"

	"./wechat"

	"github.com/gin-gonic/gin"
)

func main() {
	wechat := wechat.NewWechatWork()
	if success := wechat.Gettoken(); !success {
		log.Fatalln("Gettoken failed!")
	}

	r := gin.Default()
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	r.GET("/user/login", func(c *gin.Context) {
		code := c.Query("code")
		state := c.Query("state")
		appid := c.Query("appid")
		log.Printf("Code: %v State: %v Appid: %v\n", code, state, appid)

		UserId, errcode := wechat.Getuserinfo(code)
		if errcode == 0 {
			user, _ := wechat.Getuser(UserId)
			log.Printf("%v", user)
		}

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
