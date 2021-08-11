package models

type Address struct {
	Id       int64  `gorm:"primary"`
	UserId   int    `gorm:"Column:user_id"`
	Province string `gorm:"Column:province"`
	City     string `gorm:"Column:city"`
	Address  string `gorm:"Column:address"`
}
