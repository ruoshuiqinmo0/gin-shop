package controllers

import (
	"gin-shop/lib"
	"gin-shop/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
)

func LoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login", nil)
}

// @Summary 登录
// @Description 账号密码登录
// @Tags 用户信息   //swagger API分类标签, 同一个tag为一组
// @accept json  //浏览器可处理数据类型，浏览器默认发 Accept: */*
// @Produce  json  //设置返回数据的类型和编码
// @Param id path int true "name"    //url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
// @Param name query string false "name"
// @Success 200 {object} Res {"code":200,"data":null,"msg":""}  //成功返回的数据结构， 最后是示例
// @Failure 404 {object} Res {"code":404,"data":null,"msg":""}
// @Router /v1/login [get]    //路由信息，一定要写上
func Login(c *gin.Context) {
	var User models.User
	if err := c.ShouldBind(&User); err != nil {
		lib.Error(c, 404, err.Error())
		return
	}
	user, err := models.FindUser(User.Username, User.Password)
	if err != nil {
		lib.Error(c, 404, err.Error())
	} else {
		//c.SetCookie("user", string(user.UserId), 2*3600, "/", "localhost", false, true)
		//c.JSON(http.StatusOK, gin.H{"code":200,"msg":"success","data":user})
		token, err := lib.GenToken(user.Username)
		if err != nil {
			lib.Error(c, 404, err.Error())
		}
		lib.Success(c, token)
	}
}

func Register(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			lib.Error(c, 404, errs.Error())
			return
		}
		lib.Error(c, 404, errs.Translate(lib.Trans))
		return
	}
	id, err := models.Create(user)
	if err != nil {
		lib.Error(c, 404, err.Error())
	} else {
		lib.Success(c, id)
	}
}

func Delete(c *gin.Context) {
	idString := c.DefaultQuery("id", "0")
	id, _ := strconv.Atoi(idString)
	if id == 0 {
		lib.Error(c, 404, "id is null")
	}
	err := models.DeleteUser(id)
	if err != nil {
		lib.Error(c, 404, err.Error())
	}
	lib.Success(c, "")
}
