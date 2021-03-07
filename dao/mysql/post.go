package mysql

import (
	"strings"
	"web/bluebull/models"

	"github.com/jmoiron/sqlx"
)

func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id,author_id,community_id,title,content)
	values(?,?,?,?,?)
	`
	_, err = db.Exec(sqlStr, post.ID, post.AuthorID, post.CommunityID, post.Title, post.Content)
	return err
}

// GetPostByID 根据id查询单个帖子数据
func GetPostByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id=?
	`
	err = db.Get(post, sqlStr, id)
	return
}

// GetPostList 查询帖子列表
func GetPostList(page, size int64) (data []*models.Post, err error) {
	data = make([]*models.Post, 0, 2)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	limit ?,?
	`
	err = db.Select(&data, sqlStr, (page-1)*size, size)
	return
}

func GetPOstListByIDs(ids []string) (data []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id in (?)
	order by FIND_IN_SET(post_id,?)
	`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&data, query, args...)
	return
}
