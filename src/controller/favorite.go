package controller

import (
	"dou-xiao-yin/src/json_model"
	"dou-xiao-yin/src/service"
	"dou-xiao-yin/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteList(c *gin.Context) {

}

func FavoriteAction(c *gin.Context) {
	// 本来的逻辑是校验user_id和token的一致性
	// 但旧版客户端不能返回user_id，只能通过token找user_id。5.22新版客户端不能触发点赞 = =
	//userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	userId, _, _, _ := utils.ParseToken(token)
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	actionType, _ := strconv.Atoi(c.Query("action_type")) //1-点赞，2-取消点赞

	//fmt.Println("鉴权", userId, token)
	if userId == 0 || !service.VerifyUser(userId, token) { // 鉴权失败
		fmt.Println("token = ", token, "鉴权失败")
		c.JSON(http.StatusBadRequest, json_model.Response{StatusCode: 1, StatusMsg: "token失效"})
	} else { // 鉴权成功
		if actionType == 1 { // 点赞操作
			if err := service.FavoriteAction(videoId, userId); err != nil {
				c.JSON(http.StatusBadRequest, json_model.Response{StatusCode: 1})
			} else {
				c.JSON(http.StatusOK, json_model.Response{StatusCode: 0})
			}
		} else if actionType == 2 { // 取消点赞
			if err := service.UnFavoriteAction(videoId, userId); err != nil {
				c.JSON(http.StatusBadRequest, json_model.Response{StatusCode: 1})
			} else {
				c.JSON(http.StatusOK, json_model.Response{StatusCode: 0})
			}
		}
	}

}
