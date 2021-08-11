package controllers

import (
	"gin-shop/lib"
	"gin-shop/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	//goods := make(map[int]int,10)
	goods := c.PostForm("goods")
	userId := c.PostForm("userId")
	addressId := c.PostForm("addressId")
	if len(goods) == 0 {
		lib.Error(c, 404, "")
	}
	models.CreateOrder(goods, userId, addressId)

}
