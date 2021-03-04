package redis

import (
	"errors"
	"math"
	"time"

	"github.com/go-redis/redis"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVOte     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
)

func CreatePost(postID int64) error {
	pipeline := rdb.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	pipeline.ZAdd(getRedisKey(KeyPostScore), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})
	_, err := pipeline.Exec()
	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 去 redis 获取帖子的发布时间
	postTime := rdb.ZScore(getRedisKey(KeyPostTime), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 取当前用户给当前帖子的投票记录
	ov := rdb.ZScore(getRedisKey(KeyPostVotedPrefix+postID), userID).Val()
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value)
	pipeline := rdb.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScore), diff*op*scorePerVOte, postID)
	// 记录用户为该帖子投票的数据
	if value == 0 {
		rdb.ZRem(getRedisKey(KeyPostVotedPrefix+postID), userID)
	} else {
		rdb.ZAdd(getRedisKey(KeyPostVotedPrefix+postID), redis.Z{
			Score:  value,
			Member: userID,
		})
	}
	_, err := pipeline.Exec()
	return err
}
