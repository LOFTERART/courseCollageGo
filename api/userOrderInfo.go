package api

import (
	"PC/services/orderList"
	"github.com/gin-gonic/gin"
)

// @Tags 首页
//@Summary 小程序首页随机拼单列表和拼单科目
//@Success 20000     "成功"
// @Router /api/v1/orders [get]
func GetOrderList(c *gin.Context) {
	var lists orderList.OrderListService
	res := lists.GetOrderList()
	c.JSONP(200, &res)
}

//更新发货状态 //发货状态1 未发货 2 发货
func UpOrderSendType(c *gin.Context) {

	order := orderList.OrderListService{}

	if err := c.ShouldBind(&order); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := order.UpOrderSend()
		c.JSON(0, &res)
	}

}
