package api

import (
	service "PC/services"
	"github.com/gin-gonic/gin"
)

//增加Banner
func CreatBanner(c *gin.Context) {

	info := service.Banner{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.CreateBanner()
		c.JSON(200, &res)
	}

}

//获取Banner ALL
func GetBanner(c *gin.Context) {

	info := service.Banner{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetBanner()
		c.JSON(200, &res)
	}

}

//获取Banner ONE
func GetBannerONE(c *gin.Context) {

	info := service.Banner{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetBannerONE()
		c.JSON(200, &res)
	}

}

//更新Banner
func UpdataBanner(c *gin.Context) {

	info := service.Banner{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.UpdataBanner()
		c.JSON(200, &res)
	}

}
