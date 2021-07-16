package models

import (
	"gin-shop/dao"
	"strconv"
	"time"
)

type UserCart struct {
	CartId     int   `gorm:"primaryKey;Column:cart_id"`
	UserId     int   `gorm:"Column:user_id"`
	GoodsId    int   `gorm:"Column:goods_id"`
	CreatedAt  int64 `gorm:"Column:created_at"`
	UpdatedAt  int64 `gorm:"Column:updated_at"`
	DeleteTime int64 `gorm:"Column:delete_time"`
}

func AddCart(userIdStr, goodsIdstr string) error {
	userId, _ := strconv.Atoi(userIdStr)
	goodsId, _ := strconv.Atoi(goodsIdstr)
	var user User
	err := dao.DB.Where("user_id= ?", userId).First(&user).Error
	if err != nil {
		return err
	}
	var Goods Goods
	err = dao.DB.Where("goods_id= ?", goodsId).First(&Goods).Error
	if err != nil {
		return err
	}
	err = dao.DB.Create(&UserCart{
		UserId:     userId,
		GoodsId:    goodsId,
		CreatedAt:  0,
		UpdatedAt:  time.Now().Unix(),
		DeleteTime: 0,
	}).Error
	return err
}
