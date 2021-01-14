package mysql

import (
	"database/sql"
	"web/bluebull/models"

	"go.uber.org/zap"
)

func GetCommunityList() (data []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err = db.Select(&data, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community")
			err = nil
		}
	}
	return
}

func GetCommunityDetail(id int64) (data *models.CommunityDetail, err error) {
	data = new(models.CommunityDetail)
	sqlStr := `select 
				community_id,community_name,introduction,create_time 
				from community
				where community_id=?
			`
	if err = db.Get(data, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}
