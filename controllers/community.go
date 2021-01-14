package controllers

import (
	"database/sql"
	"strconv"
	"web/bluebull/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// 跟社区相关

func CommunityHandler(c *gin.Context) {
	data, err := logic.GetCommunityList()
	if err != nil {
		if err != sql.ErrNoRows {
			zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		}
		ResponseError(c, CodeServerBusy) //不轻易把服务端的错误暴露给外面
		return
	}
	ResponseSuccess(c, data)
}

func CommunityDetailHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	data, err := logic.GetCommunityDetail(id)
	if err != nil {
		zap.L().Error("GetCommunityDetail failed", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	ResponseSuccess(c, data)
}
