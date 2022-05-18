/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/11 19:43
 */

package main

import (
	"dou-xiao-yin/src/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化静态资源router
	router.InitResourceRouters(r)
	// 初始化接口router
	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
