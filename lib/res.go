package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data})
}

func Error(ctx *gin.Context, code int, msg string) {
	ctx.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": ""})
}
