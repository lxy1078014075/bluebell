package logic

import (
	"web/bluebull/dao/mysql"
	"web/bluebull/models"
	"web/bluebull/pkg/jwt"
	"web/bluebull/pkg/snowflake"
)

// 存放业务逻辑的代码
func SignUp(p *models.ParamSignUp) (err error) {
	// 判断用户是否存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 生成ID
	userID := snowflake.GetID()
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 构造一个user实例
	// 密码加密
	// 入库
	return mysql.InsertUser(user)
}

func Login(l *models.ParamLogin) (token string, err error) {
	user := &models.User{
		Username: l.Username,
		Password: l.Password,
	}
	if err := mysql.LoginUser(user); err != nil {
		return "", err
	}
	return jwt.GetToken(user.UserID, user.Username)
}
