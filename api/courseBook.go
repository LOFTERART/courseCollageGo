package api

import (
	"PC/services/courseBook"
	"fmt"
	"github.com/gin-gonic/gin"
)

//创建书籍
func CreateCourseBook(c *gin.Context) {
	info := courseBook.Book{}
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		fmt.Println(info, "-------------create_people----------")
		res := info.CreateCourseBook()
		c.JSON(0, &res)
	}
}

//获取书籍
func GetCourseBook(c *gin.Context) {
	info := courseBook.Book{}
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetCourseBook()
		c.JSON(0, &res)
	}
}

//删除书籍
func DelCourseBook(c *gin.Context) {
	info := courseBook.Book{}
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.DelCourseBook()
		c.JSON(0, &res)
	}
}
