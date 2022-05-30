/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package controller

import (
	"dou-xiao-yin/src/json_model"
	"dou-xiao-yin/src/service"
	"dou-xiao-yin/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	处理/douyin/feed/接口
*/
type FeedResponse struct {
	json_model.Response
	VideoList []*json_model.Video `json:"video_list,omitempty"`
	NextTime  int64               `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	latestTime := c.Query("latest_time")
	// 根据token查找用户id
	userId, _, _, _ := utils.ParseToken(token)
	videoList, nextTime := service.GetVideoList(userId, latestTime)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  json_model.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}
