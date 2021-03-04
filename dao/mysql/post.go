package mysql

import "web/bluebull/models"

func CreatePost(post *models.Post) (err error) {
	sqlStr := `insert into post(
	post_id,author_id,community_id,title,content)
	values(?,?,?,?,?)
	`
	_, err = db.Exec(sqlStr, post.ID, post.AuthorID, post.CommunityID, post.Title, post.Content)
	return err
}

func GetPostByID(id int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	where post_id=?
	`
	err = db.Get(post, sqlStr, id)
	return
}

func GetPostList() (data []*models.Post, err error) {
	data = make([]*models.Post, 0, 2)
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
	from post
	limit 10
	`
	err = db.Select(&data, sqlStr)
	return
}
