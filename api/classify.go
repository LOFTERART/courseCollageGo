package api

import (
	service "PC/services"
	"github.com/gin-gonic/gin"
)

//增加分类
func CreatClassify(c *gin.Context) {

	info := service.Classify{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.CreateClassify()
		c.JSON(200, &res)
	}

}

//获取分类
func GetClassify(c *gin.Context) {

	info := service.Classify{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetClassify()
		c.JSON(200, &res)
	}

}

//更新分类
func UpdataClassify(c *gin.Context) {

	info := service.Classify{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.UpdataClassifyS()
		c.JSON(200, &res)
	}

}

//查询分类下对应的文章
func GetClassifyArticle(c *gin.Context) {

	info := service.ClassifyArticle{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetClassifyArticle()
		c.JSON(200, &res)
	}

}
