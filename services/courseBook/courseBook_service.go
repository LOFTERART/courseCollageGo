package courseBook

import (
	"PC/models"
	"PC/serializer"
	"github.com/jinzhu/gorm"
)

type Book struct {
	Id           uint   `json:"id"`
	Title        string `json:"title"`
	Tip          string `json:"tip"`
	Img          string `json:"img"`
	CreatePeople string `json:"create_people"`
	CourseID     uint   `json:"course_id"` //课程ID
}

func (item *Book) CreateCourseBook() serializer.Response {
	book := models.Book{
		Model:        gorm.Model{},
		Title:        item.Title,
		Tip:          item.Tip,
		Img:          item.Img,
		CreatePeople: item.CreatePeople,
		CourseID:     item.CourseID,
	}
	models.DB.Create(&book)
	return serializer.Response{
		Code:  20000,
		Data:  nil,
		Msg:   "",
		Error: "",
	}
}

//传递分类和不传递分类
func (item *Book) GetCourseBook() serializer.Response {

	var book []models.Book

	if item.CourseID > 0 {
		models.DB.Preload("Course").Where("course_id=?", item.CourseID).Find(&book)
	} else {
		models.DB.Preload("Course").Find(&book)
	}

	return serializer.Response{
		Code: 20000,
		Data: serializer.FormatBookList(book),
	}

}

func (item *Book) DelCourseBook() serializer.Response {
	var book []models.Book

	models.DB.Delete(&book, item.Id)

	return serializer.Response{
		Code: 20000,
		Data: nil,
		Msg:  "删除成功",
	}

}
