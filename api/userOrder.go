package api

import (
	"PC/services/getSuccessOrder"
	"github.com/gin-gonic/gin"
)

func GetSuccessOrders(c *gin.Context) {

	info := getSuccessOrder.SuccessOrder{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetSuccessOrder()
		c.JSON(200, &res)
	}

}
