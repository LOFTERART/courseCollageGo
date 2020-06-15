package api

import (
	"PC/models"
	"PC/serializer"
	"PC/services/admin"
	"github.com/gin-gonic/gin"
)

func CreateAdminUser(c *gin.Context) {
	var userAdmin admin.Admin

	if err := c.ShouldBind(&userAdmin); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := userAdmin.CreateAdmin()
		c.JSON(200, &res)
	}

}

func LoginAdminUser(c *gin.Context) {
	var userAdmin admin.Admin

	if err := c.ShouldBind(&userAdmin); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := userAdmin.LoginAdmin()
		c.JSON(200, &res)
	}

}

func GetAdminUserInfo(c *gin.Context) {
	var userAdmin admin.Admin

	if err := c.ShouldBind(&userAdmin); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := userAdmin.GetAdmin()
		c.JSON(200, &res)
	}

}

//获取首页信息

type HomeInfo struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Num  uint   `json:"num"`
}

func GetAdminHomeInfo(c *gin.Context) {

	var (
		countUser, countOrder, countCourse, countBook uint
	)

	models.DB.Table("user").Count(&countUser)
	models.DB.Table("user_order").Where("pay_ok=?", 2).Count(&countOrder)
	models.DB.Table("course").Count(&countCourse)
	models.DB.Table("book").Count(&countBook)

	var info []*models.Course

	models.DB.Preload("UserOrder", "pay_ok=?", 2).Select("id,name,class_student,surplus_student,class_price,class_or_price").Find(&info)

	c.JSON(200, gin.H{
		"code": 20000,
		"data": gin.H{
			"infoNum": []HomeInfo{
				{1, "用户数", countUser},
				{2, "订单数", countOrder},
				{3, "课程数", countCourse},
				{4, "书籍数", countBook},
			},
			"courseNum": serializer.BuildCourseSerializers(info),
		},
	})

}
