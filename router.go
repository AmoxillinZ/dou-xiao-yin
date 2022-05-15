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

}
