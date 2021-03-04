package logic

import (
	"web/bluebull/dao/mysql"
	"web/bluebull/dao/redis"
	"web/bluebull/models"
	"web/bluebull/pkg/snowflake"

	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	p.ID = snowflake.GetID()

	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	return redis.CreatePost(p.ID)
}

func GetPostByID(id int64) (data *models.ApiPostDetail, err error) {
	post, err := mysql.GetPostByID(id)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(id) failed",
			zap.Int64("id", id),
			zap.Error(err))
		return
	}
	user, err := mysql.GetUserByID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
			zap.Int64("authorID", post.AuthorID),
			zap.Error(err))
		return
	}
	community, err := mysql.GetCommunityDetail(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetail(post.CommunityID) failed",
			zap.Int64("CommunityID", post.CommunityID),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		UserName:        user.Username,
		Post:            post,
		CommunityDetail: community,
	}
	return
}

func GetPostList() (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList()
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return
	}
	for _, post := range posts {
		user, err := mysql.GetUserByID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByID(post.AuthorID) failed",
				zap.Int64("authorID", post.AuthorID),
				zap.Error(err))
			continue
		}
		community, err := mysql.GetCommunityDetail(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetail(post.CommunityID) failed",
				zap.Int64("CommunityID", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			UserName:        user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
