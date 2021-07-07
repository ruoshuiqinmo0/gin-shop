package main

import (
	"gin-shop/dao"
	"gin-shop/routes"
)

func main(){

	err := dao.InitDb()
	if err !=nil{
		panic(err)
	}
	defer dao.Close()

	r := routes.StepRoute()

	r.Run(":8080")
}
