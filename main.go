/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/11 19:43
 */

package main

import (
	"dou-xiao-yin/config"
	"dou-xiao-yin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化配置
	//config.InitConf("./conf.yaml") //部署到服务器用这个配置，yaml文件在服务器中与可执行文件放在同一目录
	config.InitConf("./config/conf.yaml") // 开发环境配置

	// 初始化接口router
	router.InitRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
