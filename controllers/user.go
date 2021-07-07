package controllers

import (
	"gin-shop/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Login(c *gin.Context){
	var User models.User
	if err := c.ShouldBind(&User);err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	}
	user, err := models.FindUser(User.Username, User.Password)
	if err != nil{
		c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
	}else{
		c.SetCookie("user",string(user.UserId),2*3600,"/","localhost",false,true)
		c.JSON(http.StatusOK, gin.H{"user":user})
	}

}

func Register(c *gin.Context){
	var user models.User
	if err := c.ShouldBind(&user); err !=nil{
		c.JSON(http.StatusNotFound,gin.H{"err":err.Error()})
	}
	id, err:=models.Create(user)
	if err != nil{
		c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
	}else{
		c.JSON(http.StatusOK,gin.H{"id":id})
	}

}
