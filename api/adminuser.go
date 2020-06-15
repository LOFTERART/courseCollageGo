package api

import (
	"PC/services/adminuser"
	"github.com/gin-gonic/gin"
)

//获取后台用户信息
func GetAdminUser(c *gin.Context) {

	var adminUser adminuser.AdminUser

	if err := c.ShouldBind(&adminUser); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := adminUser.GetAdminUsers()
		c.JSON(0, &res)
	}

}
