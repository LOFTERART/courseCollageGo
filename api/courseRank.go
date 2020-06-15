package api

import (
	"PC/services"
	"github.com/gin-gonic/gin"
)

func GetCourseRank(c *gin.Context) {
	var service service.DailyRankService

	if err := c.ShouldBind(&service); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := service.GetCourseRank()
		c.JSON(0, &res)
	}

}
