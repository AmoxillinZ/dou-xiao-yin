/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/13 20:33
 */

package controller

import (
	"dou-xiao-yin/src/service"
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
	c.JSON(http.StatusOK, FeedResponse{
		Response:  service.Response{StatusCode: 0},
		VideoList: service.GetVideoList(),
		NextTime:  time.Now().Unix(),
	})
}
