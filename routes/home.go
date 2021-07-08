package routes

import (
	"gin-shop/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title web前台
// @version 1.0
// @description docker监控服务后端API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath
func HomeRoute(r *gin.Engine) {
	v1Group := r.Group("v1")
	{
		v1Group.GET("/login", controllers.Login)
		v1Group.GET("/register", controllers.Register)
		v1Group.GET("/delete", controllers.Delete)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
