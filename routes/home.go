package routes

import (
	"gin-shop/controllers"
	"github.com/gin-gonic/gin"
)

func HomeRoute(r *gin.Engine){
	v1Group := r.Group("v1")
	{
		v1Group.GET("/login", controllers.Login)
		v1Group.GET("/register", controllers.Register)
	}
}
