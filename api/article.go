package api

import (
	service "PC/services"
	"fmt"
	"github.com/gin-gonic/gin"
)

//添加文章
func CreateArticle(c *gin.Context) {

	info := service.Article{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		fmt.Println(info, "----------info-------")
		res := info.AddCreateArticle()
		c.JSON(200, &res)
	}

}

//获取文章多个
func GetArticle(c *gin.Context) {

	info := service.Article{}

	res := info.GetCreateArticle()
	c.JSON(200, &res)

}

//获取文章1个
func GetArticleOne(c *gin.Context) {
	info := service.Article{}
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetCreateArticleOne()
		c.JSON(200, &res)
	}

}
