package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDb()(err error){
	DB, err = gorm.Open("mysql","root:root@(127.0.0.1:3306)/db_test?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "sp_" + defaultTableName
	}
	DB.SingularTable(true)
	return DB.DB().Ping()
}

func Close(){
	err := DB.Close()
	if err != nil{
		panic(err)
	}
}
