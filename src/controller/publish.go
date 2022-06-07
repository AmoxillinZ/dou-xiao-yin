package controller

import (
	"dou-xiao-yin/src/json_model"
	"dou-xiao-yin/src/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PublishListResponse struct {
	json_model.Response
	VideoList []*json_model.Video `json:"video_list,omitempty"`
}

func PublishList(c *gin.Context) {
	// TODO 之后再补充,先写 oss 服务了
	token := c.Query("token")
	userId, _ := strconv.Atoi(c.Query("user_id"))

	//转到发布服务
	videoList, err := service.PublishList(userId, token)
	if err != nil {
		c.JSON(http.StatusBadRequest, json_model.Response{StatusCode: 1, StatusMsg: err.Error()})
	}
	c.JSON(http.StatusOK, PublishListResponse{
		Response:  json_model.Response{StatusCode: 0},
		VideoList: videoList,
	})
}

func PublishAction(c *gin.Context) {
	file, err := c.FormFile("data")
	token := c.Request.Form.Get("token")
	title := c.Request.Form.Get("title")
	if err != nil {
		c.JSON(http.StatusBadRequest, json_model.Response{StatusCode: 1, StatusMsg: err.Error()})
	}
	err = service.PublishVideo(file, token, title)
	if err != nil {
		fmt.Println(err)
		//截图可能会报错，但是功能正常，返回StatusOk
		c.JSON(http.StatusOK, json_model.Response{StatusCode: 1, StatusMsg: err.Error()})
	}
	c.JSON(http.StatusOK, json_model.Response{StatusCode: 0})
}
