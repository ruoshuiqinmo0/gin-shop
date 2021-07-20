package routes

import (
	"gin-shop/controllers"
	"gin-shop/middle"
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
	//登录 注册
	r.GET("/login", controllers.LoginIndex)
	r.POST("/login", controllers.Login)
	r.GET("/register", controllers.Register)
	v1Group := r.Group("v1")
	v1Group.Use(middle.JWT())
	{
		//商品
		v1Group.GET("/delete", controllers.Delete)
		v1Group.GET("/goodsDetail", controllers.GetGoodsDetail)
		v1Group.GET("/goodsList", controllers.GetGoodsPageList)
		//购物车
		v1Group.GET("/addCart", controllers.AddGoodsToCart)
		v1Group.GET("/delCart", controllers.DelGoodsFromCart)
		v1Group.GET("/cartList", controllers.GetCartList)
		//验证码
		v1Group.GET("/captcha", controllers.Captcha)
		v1Group.GET("/captcha/:captchaId", controllers.CaptchaImg)
		v1Group.GET("/captchas/:captchaId/:value", controllers.CaptchaVerify)

		//上传图片
		v1Group.GET("/loadImg", controllers.LoadImg)

		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
