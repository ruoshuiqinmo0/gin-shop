package models

import (
	"gin-shop/dao"
	"github.com/pkg/errors"
	"strconv"
	"time"
)

type UserCart struct {
	CartId     int   `gorm:"primaryKey;Column:cart_id"`
	UserId     int   `gorm:"Column:user_id" binding:"required" `
	GoodsId    int   `gorm:"Column:goods_id" binding:"required"`
	CreatedAt  int64 `gorm:"Column:created_at"`
	UpdatedAt  int64 `gorm:"Column:updated_at"`
	DeleteTime int64 `gorm:"Column:delete_time"`
}

type CartList struct {
	UserCart UserCart
	Goods    Goods
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

func DeleteCart(userId, goodsId int) error {
	err := dao.DB.Where("user_id = ? and goods_id = ?", userId, goodsId).Error
	return err
}

func PageList(pageStr, pageSizeStr, userIdStr string) ([]CartList, error) {
	page, _ := strconv.Atoi(pageStr)
	pageSize, _ := strconv.Atoi(pageSizeStr)
	userId, _ := strconv.Atoi(userIdStr)
	offset := (page - 1) * pageSize
	cart := make([]UserCart, pageSize)
	db := dao.DB.Where("user_id = ?", userId).Offset(offset).Limit(pageSize).Find(&cart)
	if err := db.Error; err != nil {
		return nil, err
	}
	if db.RowsAffected == 0 {
		return nil, errors.New("not Found")
	}
	list := make([]CartList, db.RowsAffected)
	for key, val := range cart {
		list[key].UserCart = val
		if err := dao.DB.Where("goods_id = ?", val.GoodsId).First(&list[key].Goods).Error; err != nil {
			return nil, err
		}
	}
	return list, nil
}
