package controllers

import (
	"fmt"
	"gin-shop/lib"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //
	ImageUrl  string `json:"imageUrl"`
}

func Captcha(c *gin.Context) {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/captcha/" + captchaId + ".png"
	c.JSON(http.StatusOK, captcha)
}

func CaptchaImg(c *gin.Context) {
	//captchaId := c.Param("captchaId")
	lib.ServerHTTP(c.Writer, c.Request)
}

func CaptchaVerify(c *gin.Context) {
	captchaId := c.Param("captchaId")

	Value := c.Param("value")
	fmt.Println(captchaId, Value)
	if captchaId == "" || Value == "" {
		lib.Error(c, 404, "请求数据不全")
	}
	if captcha.VerifyString(captchaId, Value) {
		lib.Success(c, "")
	} else {
		lib.Error(c, 404, "验证不正确")
	}

}
