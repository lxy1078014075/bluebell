package redis

import "web/bluebull/models"

func GetPostIDsInOrder(p *models.ParamPostList) ([]string, error) {
	key := getRedisKey(KeyPostTime)
	if p.Order == models.OrderScore {
		key = getRedisKey(KeyPostScore)
	}
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	return rdb.ZRevRange(key, start, end).Result()
}
