package main

import (
<<<<<<< HEAD
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
=======
	"github.com/gin-gonic/gin"
)


func main(){
	r := gin.default()
>>>>>>> 9bc18f00e6cafb9f6ac7891e3e58d4d61fbaa9e0

	r.Run(":8080")
}
