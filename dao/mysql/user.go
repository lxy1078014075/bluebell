package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"web/bluebull/models"
)

// 把每一步数据库操作封装成函数
//待logic层业务需求调用

const secret = "lxy1078014075"

func CheckUserExist(username string) (err error) {
	sqlStr := "select count(user_id) from user where username=?"
	var count int
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return
}

func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	password := encryptPassword(user.Password)
	sqlStr := "insert into user(user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlStr, user.UserID, user.Username, password)
	return
}

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func LoginUser(user *models.User) error {
	oldPassword := user.Password
	sqlStr := "select user_id,username,password from user where username=?"
	var err error
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	if err != nil {
		return err
	}
	if encryptPassword(oldPassword) != user.Password {
		return ErrorInvalidPassword
	}
	return err
}

func GetUserByID(wid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select user_id,username from user where user_id=?`
	err = db.Get(user, sqlStr, wid)
	return
}
