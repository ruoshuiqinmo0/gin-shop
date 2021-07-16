package models

import (
	"gin-shop/dao"
)

type Goods struct {
	GoodsId        int     `gorm:"Column:goods_id;primaryKey"`
	GoodsName      string  `grom:"Column:goods_name"`
	GoodsPrice     float64 `grom:"Column:goods_price"`
	GoodsNumber    int     `grom:"Column:goods_number"`
	GoodsWeight    int     `grom:"Column:goods_weight"`
	CatId          int     `grom:"Column:cat_id"`
	GoodsIntroduce string  `grom:"Column:goods_introduce"`
	GoodsBigLogo   string  `grom:"Column:goods_big_logo"`
	GoodsSmallLogo string  `grom:"goods_small_logo"`
	IsDel          int     `grom:"Column:is_del"`
	AddTime        int     `grom:"Column:add_time"`
	UpdateTime     int     `grom:"Column:update_time"`
	DeleteTime     int     `grom:"Column:delete_time"`
	CatOneId       int     `grom:"Column:cat_one_id"`
	CatTwoId       int     `grom:"Column:cat_two_id"`
	CatThreeId     int     `grom:"Column:cat_three_id"`
	HotNumber      int     `grom:"Column:hot_number"`
	IsPromote      int     `grom:"Column:is_promote"`
	GoodsState     int     `grom:"Column:goods_name"`
}

type GoodsAttr struct {
	ID        int    `gorm:"Column:id"`
	GoodsId   int    `gorm:"Column:goods_id"`
	AttrId    int    `gorm:"Column:attr_id"`
	AttrValue string `gorm:"Column:attr_value"`
}

type GoodsDetail struct {
	Goods     Goods
	GoodsAttr []GoodsAttr
}

func GetGoodsDetail(id int) (goodsDetail GoodsDetail, err error) {
	var goods Goods
	err = dao.DB.Where("goods_id = ?", id).Find(&goods).Error
	if err != nil {
		return goodsDetail, err
	}
	var attr []GoodsAttr
	err = dao.DB.Where("goods_id = ?", id).Find(&attr).Error
	if err != nil {
		return goodsDetail, err
	}
	return GoodsDetail{
		Goods:     goods,
		GoodsAttr: attr,
	}, nil
}

func GetListGoods(page int, pageSize int) ([]GoodsDetail, error) {
	offset := (page - 1) * pageSize
	var goods []Goods
	err := dao.DB.Offset(offset).Limit(pageSize).Find(&goods).Error
	if err != nil {
		return nil, err
	}
	list := make([]GoodsDetail, pageSize)
	for key, val := range goods {
		list[key] = GoodsDetail{
			Goods: val,
		}
		err := dao.DB.Where("goods_id = ?", val.GoodsId).Find(&(list[key].GoodsAttr)).Error
		if err != nil {
			return nil, err
		}
	}
	return list, err
}
