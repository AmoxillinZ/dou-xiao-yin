/**
 * @Author: Amo
 * @Description:
 * @Date: 2022/5/26 11:18
 */

package controller

import (
	"dou-xiao-yin/src/service"
	"dou-xiao-yin/src/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	token := c.Query("token")
	toUserId, _ := strconv.Atoi(c.Query("to_user_id"))
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	var err error
	// TODO 客户端不能返回user_id,先用token查出来
	userId, _ = utils.GetIdFromToken(token)
	if err = service.TokenVerify(userId, token); err != nil { // 鉴权失败
		c.JSON(http.StatusOK, FeedResponse{
			Response: service.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}
	// 调用关注/取关方法
	err = service.RelationAction(userId, toUserId, actionType)
	if err != nil { // 操作失败
		c.JSON(http.StatusOK, FeedResponse{
			Response: service.Response{StatusCode: 2, StatusMsg: err.Error()},
		})
		return
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: service.Response{StatusCode: 0},
	})
	return
}
