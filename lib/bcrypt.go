package lib

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

)


func Validate(src string, password string)(bool,error){
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(src));err !=nil{
		return false, errors.New("匹配")
	}
	return true, nil
}

func Generate(src string)([]byte, error){
	return bcrypt.GenerateFromPassword([]byte(src), bcrypt.DefaultCost)
}