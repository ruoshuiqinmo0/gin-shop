package routes

import (
	"gin-shop/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoginCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//Path := c.Request.URL.Path
		//method := c.Request.Method
		//if (Path == "/v1/login" || Path == "/v1/register") && method == "GET"{
		//
		//}else {
		//	user, err := c.Cookie("user")
		//	if err != nil {
		//		c.Redirect(http.StatusMovedPermanently, "/v1/login")
		//	}
		//	userID, err := strconv.Atoi(user)
		//	if err != nil || userID <= 0 {
		//		c.Redirect(http.StatusMovedPermanently, "/v1/login")
		//	}
		//}
	}

}

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
	r.LoadHTMLGlob("views/Home/**/*")
	v1Group := r.Group("v1")
	v1Group.Use(LoginCheck())
	{
		v1Group.GET("/login", controllers.LoginIndex)
		v1Group.POST("/login", controllers.Login)
		v1Group.GET("/register", controllers.Register)
		v1Group.GET("/delete", controllers.Delete)
		v1Group.GET("/goodsDetail", controllers.GetGoodsDetail)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
