/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package controller

import (
	"dou-xiao-yin/src/service"
	"dou-xiao-yin/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
	处理/douyin/feed/接口
*/
type FeedResponse struct {
	service.Response
	VideoList []*service.Video `json:"video_list,omitempty"`
	NextTime  int64            `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {
	token := c.Query("token")
	//lastTime := c.Query("last_time")
	// TODO 需要完善last_time参数
	// 根据token查找用户id
	userId, _, _, _ := utils.ParseToken(token)
	c.JSON(http.StatusOK, FeedResponse{
		Response:  service.Response{StatusCode: 0},
		VideoList: service.GetVideoList(userId),
		// TODO nextTime:的值需要完善
		NextTime: time.Now().Unix(),
	})
}
