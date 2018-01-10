package main

import (
"github.com/gin-gonic/gin"
"net/http"
"fmt"
"github.com/astaxie/beego"
)

func getQuery(context *gin.Context)  {
	userid := context.Query("userid")
	username := context.Query("username")

	context.String(http.StatusOK,userid + "" + username)
}
func getParam(context *gin.Context)  {
	userid := context.Param("userid")
	username := context.Param("username")

	context.String(http.StatusOK,userid + "" + username)
}
func post(context *gin.Context)  {
	username := context.PostForm("username")
	password := context.DefaultPostForm("password","123456")
	context.JSON(200,gin.H{
		"status":"posted",
		"username":username,
		"password":password,
	})
}

type User struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

func bind(context *gin.Context) {
	var user User
	if context.Bind(&user) == nil {
		context.JSON(http.StatusOK, gin.H{
			"username": user.UserName,
			"password": user.PassWord,
		})
	}
}

func nick(c *gin.Context)  {
	message := c.PostForm("message")
	nick := c.DefaultPostForm("nick","anonymous")

	c.JSON(200,gin.H{
		"status":"posted",
		"message":message,
		"nick":nick,
	})
}

func popo(c *gin.Context)  {
	id := c.Query("id")
	page := c.DefaultQuery("page","0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id:%s; page:%s; name:%s; message:%s",id,page,name,message)

}

func main() {
	router := gin.Default()
	router.GET("/user",getQuery)
	router.GET("user/:userid/:username",getParam)
	router.POST("/pos",post)
	router.POST("/bind",bind)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"pong",
		})
	})
	router.POST("/formpost",nick)
	router.POST("post",popo)
	router.Run(":8000")
}

type LoginContrller struct {
	beego.Controller
}

func (c *LoginContrller) Get() {
	c.TplName = "login.html"
}








