package models

import (
	"errors"
	"gin-shop/dao"
	"gin-shop/lib"
)

type User struct {
	UserId   int    `gorm:"primary_key"`
	Username string `gorm:"Column:username" form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func FindUser(name string, password string) (user User, err error) {
	//user = new(LoginUser)
	err = dao.DB.Where("username = ?", name).First(&user).Error
	if err != nil {
		return user, err
	}
	isOk, _ := lib.Validate(password, user.Password)
	if !isOk {
		return user, errors.New("密码错误")
	}
	return user, nil
}

func Create(user User) (id int, err error) {
	password, _ := lib.Generate(user.Password)
	user.Password = string(password)
	db := dao.DB.FirstOrCreate(&user, User{Username: user.Username})
	if db.Error != nil {
		return 0, err
	} else if db.RowsAffected == 0 {
		return 0, errors.New("已存在了")
	}
	return user.UserId, nil
}

func DeleteUser(id int) error {
	return dao.DB.Where("user_id = ?", id).Delete(&User{}).Error
}
