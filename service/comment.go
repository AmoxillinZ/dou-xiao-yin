package service

import (
	"dou-xiao-yin/json_model"
	mapper2 "dou-xiao-yin/model"
)

func AddComment(userId int, videoId int, commentText string) (json_model.Comment, error) {

	commentRes := json_model.Comment{}
	// 新增评论
	comment := &mapper2.Comment{VideoId: videoId, UserId: userId, Content: commentText}
	commentId, err := mapper2.AddComment(comment)
	if err != nil {
		return commentRes, err
	}

	// 增加视频评论数
	if err := mapper2.AddCommentCount(videoId); err != nil {
		return commentRes, err
	}

	// 获取数据库中 comment
	comment, err = mapper2.GetCommentById(commentId)
	if err != nil {
		return commentRes, err
	}

	// 获取视频 authorId
	authorId, err := mapper2.GetAuthorIdByVideoId(videoId)
	if err != nil {
		return commentRes, err
	}

	// 获取当前评论用户的数据
	user, err := mapper2.GetUserById(userId)
	if err != nil {
		return commentRes, err
	}

	// 创建 response 中的 user 结构
	userRes := &json_model.User{
		Id:            userId,
		Username:      user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      mapper2.IsFollow(userId, authorId)}

	// 创建 response 中的 comment 结构
	createAt := comment.CreateTime.Time.Format("01-02") // 评论发布日期 mm-dd
	commentRes = json_model.Comment{Id: commentId, User: *userRes, Content: commentText, CreateDate: createAt}

	return commentRes, nil
}

func DeleteComment(userId int, videoId int, commentId int) error {

	comment := &mapper2.Comment{Id: commentId, VideoId: videoId, UserId: userId}
	if err := mapper2.DeleteComment(comment); err != nil {
		return err
	}

	// 减少视频评论数
	if err := mapper2.DeleteCommentCount(videoId); err != nil {
		return err
	}

	return nil
}

func GetComments(videoId int) ([]*json_model.Comment, error) {

	// 评论 response
	commentsRes := make([]*json_model.Comment, 0)

	// 倒序获取数据库中对应视频下所有 comment
	comments, err := mapper2.GetCommentsByVideoId(videoId)
	if err != nil {
		return nil, err
	}

	// 当前视频作者 id
	authorId, err := mapper2.GetAuthorIdByVideoId(videoId)
	if err != nil {
		return nil, err
	}

	// 循环处理每个 comment，返回相应格式
	for _, comment := range comments {
		// 当前评论的 user 信息
		user, err := mapper2.GetUserById(comment.UserId)
		if err != nil {
			return nil, err
		}

		// 创建 response 中的 user 结构
		userRes := &json_model.User{
			Id:            user.Id,
			Username:      user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      mapper2.IsFollow(user.Id, authorId),
		}

		createAt := comment.CreateTime.Time.Format("01-02") // 评论发布日期 mm-dd
		commentRes := &json_model.Comment{Id: comment.Id, User: *userRes, Content: comment.Content, CreateDate: createAt}

		commentsRes = append(commentsRes, commentRes)
	}

	return commentsRes, nil
}
