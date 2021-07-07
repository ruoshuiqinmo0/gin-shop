package routes

import "github.com/gin-gonic/gin"

func StepRoute()(e *gin.Engine){
	gin.ForceConsoleColor()
	r := gin.Default()
	HomeRoute(r)
	AdminRoute(r)
	return r
}
