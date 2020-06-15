package api

import (
	service "PC/services"
	"github.com/gin-gonic/gin"
)

//增加studio
func CreatStudio(c *gin.Context) {

	info := service.Studio{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.CreateStudio()
		c.JSON(200, &res)
	}

}

//获取studio ALL
func GetStudio(c *gin.Context) {

	info := service.Studio{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetStudio()
		c.JSON(200, &res)
	}

}

//获取studio ONE
func GetStudioOne(c *gin.Context) {

	info := service.Studio{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetStudioOne()
		c.JSON(200, &res)
	}

}

//更新studio
func UpdataStudio(c *gin.Context) {

	info := service.Studio{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.UpdataStudio()
		c.JSON(200, &res)
	}

}
