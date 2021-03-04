package router

import (
	"net/http"
	"web/bluebull/logger"
	"web/bluebull/middlewares"

	"go.uber.org/zap"

	"web/bluebull/controllers"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {

	if err := controllers.InitTrans("zh"); err != nil {
		zap.L().Error("init validator trans failed", zap.Error(err))
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	v1 := r.Group("/api/v1")

	// 注册路由业务
	v1.POST("/signup", controllers.SignUpHandler)
	// 登陆
	v1.POST("/login", controllers.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) //应用jwt中间件
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)
		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		v1.GET("/posts", controllers.GetPostListHandler)

		v1.POST("/vote", controllers.PostVoteHandler)
	}

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK,
			gin.H{"msg": "ok",
				"userID": c.GetInt64(controllers.CtxUserIDKey),
				"air":    "ok1211",
			})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": 404,
		})
	})
	return r
}
