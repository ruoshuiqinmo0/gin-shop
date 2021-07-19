package main

import (
	"gin-shop/dao"
	"gin-shop/lib"
	"gin-shop/routes"
)

func main() {
	if err := lib.InitTrans("zh"); err != nil {
		panic(err)
		return
	}
	err := dao.InitDb()
	if err != nil {
		panic(err)
	}
	defer dao.Close()

	r := routes.StepRoute()
	r.Run(":8080")
}
