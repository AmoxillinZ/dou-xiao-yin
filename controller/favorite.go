package controller

import (
	"dou-xiao-yin/json_model"
	"dou-xiao-yin/service"
	"dou-xiao-yin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavoriteListResponse struct {
	json_model.Response
	FavoriteList []*json_model.Video `json:"video_list,omitempty"`
}

func FavoriteList(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	loginId, verifyErr := utils.GetIdFromToken(token)
	if verifyErr != nil {
		c.JSON(http.StatusBadRequest, FavoriteListResponse{
			Response: json_model.Response{StatusCode: 1},
		})
	}

	favoriteList, err := service.FavoriteList(userId, loginId)
	if err != nil {
		c.JSON(http.StatusBadRequest, FavoriteListResponse{
			Response: json_model.Response{StatusCode: 1},
		})
		return
	}

	c.JSON(http.StatusOK, FavoriteListResponse{
		Response:     json_model.Response{StatusCode: 0},
		FavoriteList: favoriteList,
	})
	return
}

func FavoriteAction(c *gin.Context) {
	// 本来的逻辑是校验user_id和token的一致性，但客户端不能返回user_id
	//userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	userId, _ := utils.GetIdFromToken(token)
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
			}
			c.JSON(http.StatusOK, json_model.Response{StatusCode: 0})
		}
		if actionType == 2 { // 取消点赞
			if err := service.UnFavoriteAction(videoId, userId); err != nil {
				c.JSON(http.StatusBadRequest, json_model.Response{StatusCode: 1})
			}
			c.JSON(http.StatusOK, json_model.Response{StatusCode: 0})
		}
	}

}
