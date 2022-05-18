/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/11 19:43
 */

package router

import (
	"dou-xiao-yin/src/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./public/")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller.Feed)

}
