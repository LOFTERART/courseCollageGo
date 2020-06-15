package api

import (
	"PC/services/payCourseOrder"
	"github.com/gin-gonic/gin"
)

func PayCourseOrder(c *gin.Context) {

	payOrder := payCourseOrder.PayCourseOrder{}

	if err := c.ShouldBind(&payOrder); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := payOrder.PayOrder()
		c.JSON(0, &res)
	}

}
