package payCourseOrder

import (
	"PC/models"
	"PC/serializer"
	"github.com/jinzhu/gorm"
)

type PayCourseOrder struct {
	OrderID  uint `json:"orderId"`
	CourseId uint `json:"courseId"` //课程id 用来减少课程学生数量
}

func (pay *PayCourseOrder) PayOrder() serializer.Response {

	order := models.UserOrder{
		Model: gorm.Model{
			ID: pay.OrderID,
		},
	}

	models.DB.Model(&order).Update("pay_ok", 2)

	course := models.Course{
		Model: gorm.Model{
			ID: pay.CourseId,
		},
	}

	models.DB.Model(&course).UpdateColumn("surplus_student", gorm.Expr("surplus_student - ?", 1)).First(&course)

	return serializer.Response{
		Code: 0,
		Data: map[string]interface{}{
			"pay": true,
			"stu": course.SurplusStudent,
		},
		Msg:   "",
		Error: "",
	}
}
