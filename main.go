package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"./wechat"

	"github.com/gin-gonic/gin"
)

func main() {
	domain := flag.String("domain", "127.0.0.1", "App domain")
	port := flag.Int("port", 9527, "App domain")
	flag.Parse()

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
			c.SetCookie("wcwl_uid", UserId, 7200, "/", *domain, false, true)
		}

		redirect := fmt.Sprintf("http://%s:%d%s", *domain, *port, "/")
		log.Printf("Redirecting: %s ...", redirect)
		script := fmt.Sprintf("<script>window.location.href = '%s';</script>", redirect)
		c.Header("Content-Type", "text/html")
		c.String(http.StatusOK, script)
	})
	r.Run()
}
