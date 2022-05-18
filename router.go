/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/11 19:43
 */

package main

import (
	"MyProject/src/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public/")

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/file/", func(c *gin.Context) {
		c.File("./src/public/img.png")
	})
	apiRouter.GET("/video1/", func(c *gin.Context) {
		c.File("./src/public/bear.mp4")
	})
	apiRouter.GET("/video2/", func(c *gin.Context) {
		c.File("./src/public/bear.mp4")
	})

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)

	// 用户路由组
	user := apiRouter.Group("/user")
	{
		user.GET("/", controller.UserInfo)     // 用户信息接口
		user.POST("/login/", controller.Login) // 用户登录接口
		user.POST("/register/")                // 用户注册接口
	}
	// 喜欢路由组
	favorite := apiRouter.Group("/favorite")
	{
		favorite.GET("/list/", controller.FavoriteList)
	}

	// 发布路由组
	publish := apiRouter.Group("/publish")
	{
		publish.GET("/list/", controller.PublishList)
	}

}
