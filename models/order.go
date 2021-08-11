package models

import (
	"gin-shop/dao"
	"time"
)

type Order struct {
	OrderId            int
	UserId             int
	OrderNumber        int
	OrderPrice         float64
	OrderPay           int
	IsSend             string
	TradeNo            string
	OrderFapiaoTitle   string
	OrderFapiaoCompany string
	OrderFapiaoContent string
	ConsigneeAddr      string
	PayStatus          string
	CreateTime         int64
	UpdateTime         int64
}

type OrderGoods struct {
	Id              int
	OrderId         int
	GoodsId         int
	GoodsPrice      float64
	GoodsNumber     int
	GoodsTotalPrice float64
}

func CreateOrder(postGoods map[int]int, userId int, addressId int) (orderNo string, err error) {
	var (
		Goods   Goods
		User    User
		Order   Order
		Address Address
	)
	Order.CreateTime = time.Now().Unix()
	dao.DB.Where("id = ?", userId).First(&User)
	if err != nil {
		return "", err
	}
	dao.DB.Where("id = ?", addressId).First(&Address)
	if err != nil {
		return "", err
	}
	orderPrice := 0.00
	for goodsId, number := range postGoods {
		err = dao.DB.Where("id = ?", goodsId).First(&Goods).Error
		if err != nil {
			return "", err
		}

		dao.DB.Create(&OrderGoods{
			GoodsId:         goodsId,
			GoodsPrice:      Goods.GoodsPrice,
			GoodsNumber:     number,
			GoodsTotalPrice: float64(number) * Goods.GoodsPrice,
		})
		orderPrice = orderPrice + float64(number)*Goods.GoodsPrice
	}
	Order.UserId = userId
	Order.OrderNumber = userId
	Order.OrderPrice = orderPrice
	Order.OrderPay = 0
	Order.IsSend = "否"
	Order.ConsigneeAddr = Address.Province + Address.City + Address.Address

}
