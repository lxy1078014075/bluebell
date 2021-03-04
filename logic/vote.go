package logic

import (
	"strconv"
	"web/bluebull/dao/redis"
	"web/bluebull/models"

	"go.uber.org/zap"
)

func PostVote(userID int64, p *models.ParamVote) error {
	zap.L().Debug("PostVote",
		zap.Int64("userID", userID),
		zap.String("postID", p.PostID),
		zap.Int8("direction", p.Direction),
	)
	return redis.VoteForPost(strconv.Itoa(int(userID)), p.PostID, float64(p.Direction))
}
