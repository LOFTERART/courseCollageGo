package api

import (
	"PC/services/address"
	"PC/services/userInfo"
	"github.com/gin-gonic/gin"
)

//添加地址
func GetAddress(c *gin.Context) {

	var uInfo userInfo.UserAddress

	if err := c.ShouldBind(&uInfo); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := uInfo.UpdateAddress()
		c.JSON(0, &res)
	}

}

//获取地址列表
func GetAddressList(c *gin.Context) {
	var addr address.Address
	if err := c.ShouldBind(&addr); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := addr.GetAddressLists()
		c.JSON(0, &res)
	}
}

func UpOrderAddress(c *gin.Context) {
	var addr address.UserChoseAddr
	if err := c.ShouldBind(&addr); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := addr.UpdataOrderAddress()
		c.JSON(0, &res)
	}
}
