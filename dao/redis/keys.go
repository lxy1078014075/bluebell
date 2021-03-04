package redis

const (
	KeyPrefix          = "bluebell:"
	KeyPostTime        = "post:time"  // zset;帖子与时间
	KeyPostScore       = "post:score" // zset;帖子与投票的分数
	KeyPostVotedPrefix = "post:vote:" // zset;记录用户和投票类型，参数是post_id
)

func getRedisKey(key string) string {
	return KeyPrefix + key
}
