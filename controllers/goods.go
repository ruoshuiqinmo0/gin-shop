package controllers

import (
	"gin-shop/lib"
	"gin-shop/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetGoodsDetail(c *gin.Context) {
	idStr := c.DefaultQuery("id", "0")
	if idStr == "0" {
		lib.Error(c, 404, "产品id没找到")
	}
	id, _ := strconv.Atoi(idStr)
	Goods, err := models.GetGoodsDetail(id)
	if err != nil {
		lib.Error(c, 404, err.Error())
	} else {
		lib.Success(c, Goods)
	}
}

func GetGoodsPageList(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	pageSizeStr := c.DefaultQuery("page_size", "15")
	pageSize, _ := strconv.Atoi(pageSizeStr)
	list, err := models.GetListGoods(page, pageSize)
	if err != nil {
		lib.Error(c, 404, err.Error())
	} else {
		lib.Success(c, list)
	}
}
