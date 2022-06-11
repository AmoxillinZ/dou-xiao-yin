/**
 * @Author: Amo
 * @Description: 在视频存在本地时用于为路径指定静态资源，现在静态资源放在云服务器，用不到了
 * @Date: 2022/5/17 17:47
 */

package router

import (
	"dou-xiao-yin/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

// InitResourceRouters : 为视频播放路径和封面路径指定对应资源/*
func InitResourceRouters(r *gin.Engine) error {
	resourcesPath := config.GetConf().Resource.StaticPath
	r.GET("/douyin/resources/video/:user_id/:id", func(c *gin.Context) {
		userId := c.Param("user_id")
		id := c.Param("id")
		videoPath := fmt.Sprintf("%s/videos/%s/%s.mp4", resourcesPath, userId, id)
		c.File(videoPath)
	})
	r.GET("/douyin/resources/cover/:user_id/:id", func(c *gin.Context) {
		userId := c.Param("user_id")
		id := c.Param("id")
		coverPath := fmt.Sprintf("%s/covers/%s/%s.jpg", resourcesPath, userId, id)
		c.File(coverPath)
	})
	return nil
}
