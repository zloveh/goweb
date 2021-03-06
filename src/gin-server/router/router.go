package router

import (
	"ginweb/src/controller"
	"ginweb/src/middleware"
	"github.com/gin-gonic/gin"
)

var RouterMux *gin.Engine

func Router() {
	// 默认debug, 线上需要设置成 release模式
	// gin.SetMode(gin.ReleaseMode)

	// 新建一个没有任何默认中间件的路由
	RouterMux = gin.New()

	// 加载中间件
	RouterMux.Use(gin.Logger())

	// 加入自定义中间件
	RouterMux.GET("/", middleware.Middleware1, controller.A1)
	RouterMux.GET("/welcome/:name", controller.GetParam)
	RouterMux.GET("/welcome", controller.GetQuery)
	RouterMux.POST("/post", controller.GetPostValue)
	RouterMux.POST("/build", controller.BuildData)

	// 路由组
	v1 := RouterMux.Group("/v1", middleware.Middleware1)
	{
		v1.GET("/", controller.A1)
		v1.GET("/welcome/:name", controller.GetParam)
	}
}
