package controllers

import (
	"gin-shop/lib"
	"gin-shop/models"
	"github.com/gin-gonic/gin"
)

func AddGoodsToCart(c *gin.Context) {
	userId := c.DefaultQuery("user_id", "0")
	GoodsId := c.DefaultQuery("goods_id", "0")
	err := models.AddCart(userId, GoodsId)
	if err != nil {
		lib.Error(c, 404, err.Error())
	} else {
		lib.Success(c, nil)
	}
}

func GetCartList(c *gin.Context) {

}

func DelGoodsFromCart(c *gin.Context) {
	var cart models.UserCart
	if err := c.ShouldBind(&cart); err != nil {
		lib.Error(c, 404, err.Error())
	} else {
		err := models.DeleteCart(cart.UserId, cart.GoodsId)
		if err != nil {
			lib.Error(c, 404, err.Error())
		} else {
			lib.Success(c, "")
		}
	}

}
