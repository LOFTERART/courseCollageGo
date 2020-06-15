package api

import (
	"PC/services/classInfo"
	"github.com/gin-gonic/gin"
)

//创建订单
func GetClassInfo(c *gin.Context) {
	info := classInfo.ClassInfo{}
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetClassInfo()
		c.JSON(0, &res)
	}

}

//生成微信订单 返回前端需要支付的数据
func CreateWeChatOrder(c *gin.Context) {

	weChatOrder := classInfo.WeChatOrder{}

	if err := c.ShouldBind(&weChatOrder); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := weChatOrder.CreateWeChatOrder(c.ClientIP())
		c.JSON(200, &res)
	}

}
