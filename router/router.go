/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/11 19:43
 */

package router

import (
	controller2 "dou-xiao-yin/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Static("/static", "./public/")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", controller2.Feed)

	// 用户路由组
	user := apiRouter.Group("/user")
	{
		user.GET("/", controller2.UserInfo)           // 用户信息接口
		user.POST("/login/", controller2.Login)       // 用户登录接口
		user.POST("/register/", controller2.Register) // 用户注册接口
	}
	// 喜欢路由组
	favorite := apiRouter.Group("/favorite")
	{
		favorite.GET("/list/", controller2.FavoriteList)
		favorite.POST("/action/", controller2.FavoriteAction)
	}

	// 发布路由组
	publish := apiRouter.Group("/publish")
	{
		publish.GET("/list/", controller2.PublishList)
		publish.POST("/action/", controller2.PublishAction)
	}

	// 关注路由组
	relation := apiRouter.Group("/relation")
	{
		relation.POST("/action/", controller2.RelationAction)
		relation.GET("/follow/list/", controller2.FollowList)
		relation.GET("/follower/list", controller2.FollowerList)
	}

	// 评论路由组
	comment := apiRouter.Group("/comment")
	{
		comment.GET("/list/", controller2.CommentList)
		comment.POST("/action/", controller2.CommentAction)
	}
}
