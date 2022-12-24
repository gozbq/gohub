package bootstrap

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/middlewares"
	"gohub/routes"
	"net/http"
	"strings"
)

// SetupRoute 路由初始化
func SetupRoute(router *gin.Engine) {
	// 注册全局中间件
	registerGlobalMiddleWare(router)
	//注册API路由
	routes.RegisterAPIRoutes(router)
	//配置404路由
	setup404Handler(router)
}
func registerGlobalMiddleWare(router *gin.Engine) {
	router.Use(middlewares.Logger(), middlewares.Recovery())
}
func setup404Handler(router *gin.Engine) {
	//处理404请求
	router.NoRoute(func(c *gin.Context) {
		//获取标头信息的accept信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是HTML的话
			c.String(http.StatusNotFound, "页面返回404")
		} else {
			// 默认返回JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义,请确认url和请求方法是否正确",
			})
		}
	})
}
