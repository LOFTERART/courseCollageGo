package api

import (
	"PC/services/userInfo"
	"github.com/gin-gonic/gin"
)

//获取用户信息
func GetUserInfo(c *gin.Context) {

	var u userInfo.UserInfo

	if err := c.ShouldBind(&u); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := u.AddUser(c.ClientIP())
		c.JSON(200, &res)
	}

}

//更新用户信息
func UpdataUserInfo(c *gin.Context) {

	var u userInfo.UserSubmit

	if err := c.ShouldBind(&u); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := u.UpdataSubmitCode()
		c.JSON(200, &res)
	}

}
