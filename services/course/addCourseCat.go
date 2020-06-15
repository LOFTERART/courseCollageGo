package course

import (
	"PC/models"
	"PC/serializer"
	"fmt"
)

//分类struct
type CourseCat struct {
	//id存在更新分类
	ID uint `json:"id"`
	//创建分类
	Name           string `json:"name"`
	Mark           string `json:"mark"`
	CreatePeople   string `json:"create_people"`
	SubName        string `json:"sub_name"` //课程名称
	ClassPrice     uint   `json:"class_price"`
	ClassOrPrice   uint   `json:"class_orPrice"`
	ClassTime      int64  `json:"class_time"`
	ClassEndTime   int64  `json:"class_end_time"`
	ClassStudent   uint   `json:"class_student"`
	SurplusStudent uint   `json:"surplus_student"` //剩余学生人数
	SortNum        uint   `json:"sort_num"`        //排序

	Content string `json:"content"` //文本内容
}

//创建分类
func (item *CourseCat) CreateCourseCat() serializer.Response {

	if item.ID > 0 {
		course := models.Course{
			Name:           item.Name,
			Mark:           item.Mark,
			CreatePeople:   item.CreatePeople,
			SubName:        item.SubName,
			ClassPrice:     item.ClassPrice,
			ClassOrPrice:   item.ClassOrPrice,
			ClassStudent:   item.ClassStudent,
			SurplusStudent: item.SurplusStudent,
			ClassTime:      item.ClassTime,
			ClassEndTime:   item.ClassEndTime,
			SortNum:        item.SortNum,
			Content:        item.Content,
		}
		models.DB.Where("id=?", item.ID).Model(&course).Updates(&course)

		return serializer.Response{
			Code: 20000,
			Data: serializer.BuildCourseSerializer(course),
			Msg:  "更新成功",
		}
	} else {
		course := models.Course{
			Name:           item.Name,
			Mark:           item.Mark,
			CreatePeople:   item.CreatePeople,
			SubName:        item.SubName,
			ClassPrice:     item.ClassPrice,
			ClassOrPrice:   item.ClassOrPrice,
			ClassStudent:   item.ClassStudent,
			SurplusStudent: item.SurplusStudent,
			ClassTime:      item.ClassTime,
			ClassEndTime:   item.ClassEndTime,
			SortNum:        item.SortNum,
			Content:        item.Content,
		}
		models.DB.Create(&course)

		return serializer.Response{
			Code: 20000,
			Data: serializer.BuildCourseSerializer(course),
		}

	}

}

//获取分类

func (item *CourseCat) GetCourse() serializer.Response {

	var course []*models.Course
	models.DB.Preload("Book").Order("sort_num desc").Find(&course)

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildCourseSerializers(course),
	}

}

//课程详细

type CourseDetail struct {
	ID uint `json:"id"` //课程id
}

func (item *CourseDetail) GetCourseDetail() serializer.Response {

	var course models.Course
	models.DB.First(&course, item.ID)

	fmt.Println(item.ID, "-------item-----")

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildCourseSerializer(course),
	}

}
