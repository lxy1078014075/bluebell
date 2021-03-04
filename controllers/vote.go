package controllers

import (
	"web/bluebull/logic"
	"web/bluebull/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func PostVoteHandler(c *gin.Context) {
	// 参数校验
	p := new(models.ParamVote)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("PostVote with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors) // 类型断言
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 获取用户id
	userId, err := getCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		return
	}
	err = logic.PostVote(userId, p)
	if err != nil {
		zap.L().Error("logic.PostVote(userId, p) failed", zap.Error(err))
		return
	}
	ResponseSuccess(c, nil)
}
