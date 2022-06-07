package controller

import (
	"dou-xiao-yin/src/json_model"
	"dou-xiao-yin/src/service"
	"dou-xiao-yin/src/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CommentListResponse struct {
	json_model.Response
	CommentList []*json_model.Comment `json:"comment_list,omitempty"`
}

type CommentResponse struct {
	json_model.Response
	Comment *json_model.Comment `json:"comment,omitempty"`
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	// 用户鉴权
	if err := utils.AuthenticateToken(token); err != nil {
		err := errors.New("用户鉴权失败，无法拉取评论")
		c.JSON(http.StatusBadRequest, CommentListResponse{
			Response: json_model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
	}

	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		err := errors.New("获取 video id 失败")
		c.JSON(http.StatusBadRequest, CommentListResponse{
			Response: json_model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
	}

	comments, err := service.GetComments(videoId)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommentListResponse{
			Response: json_model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
	}

	for i, comment := range comments {
		fmt.Println(i)
		fmt.Println(comment.Content)
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    json_model.Response{StatusCode: 0},
		CommentList: comments,
	})

}

func CommentAction(c *gin.Context) {
	// 用户鉴权
	token := c.Query("token")
	userId, err := utils.GetIdFromToken(token)
	if err != nil {
		err := errors.New("用户鉴权失败，无法评论")
		c.JSON(http.StatusBadRequest, CommentResponse{
			Response: json_model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
	}

	// 获取相关信息
	videoId, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		err := errors.New("获取 video id 失败")
		c.JSON(http.StatusBadRequest, CommentResponse{
			Response: json_model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
	}

	actionType, err := strconv.Atoi(c.Query("action_type"))
	if err != nil {
		err := errors.New("获取 action type 失败")
		c.JSON(http.StatusBadRequest, CommentResponse{
			Response: json_model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
	}
	commentText := c.Query("comment_text")

	commentId, err := strconv.Atoi(c.Query("comment_id"))
	if err != nil {
		err := errors.New("获取 comment id 失败")
		c.JSON(http.StatusBadRequest, CommentResponse{
			Response: json_model.Response{
				StatusCode: 1,
				StatusMsg:  err.Error()},
		})
	}

	// 发布评论
	if actionType == 1 {
		comment, err := service.AddComment(userId, videoId, commentText)
		if err != nil {
			c.JSON(http.StatusBadRequest, CommentResponse{
				Response: json_model.Response{
					StatusCode: 1,
					StatusMsg:  err.Error()},
			})
		}

		c.JSON(http.StatusOK, CommentResponse{
			Response: json_model.Response{StatusCode: 0, StatusMsg: "评论成功"},
			Comment:  comment,
		})
	}

	// 删除评论
	if actionType == 2 {
		err := service.DeleteComment(userId, videoId, commentId)
		if err != nil {
			c.JSON(http.StatusBadRequest, CommentResponse{
				Response: json_model.Response{
					StatusCode: 1,
					StatusMsg:  err.Error()},
			})
		}
		c.JSON(http.StatusOK, CommentResponse{
			Response: json_model.Response{StatusCode: 0, StatusMsg: "删除评论成功"},
		})
	}
}
