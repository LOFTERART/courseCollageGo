package api

import (
	"PC/services/course"
	"github.com/gin-gonic/gin"
)

//创建分类
func AddCatControl(c *gin.Context) {

	info := course.CourseCat{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.CreateCourseCat()
		c.JSON(200, &res)
	}
}

//获取分类

func GetCourseCat(c *gin.Context) {

	info := course.CourseCat{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetCourse()
		c.JSON(200, &res)
	}

}

func GetCourseDetail(c *gin.Context) {
	info := course.CourseDetail{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetCourseDetail()
		c.JSON(200, &res)
	}
}
