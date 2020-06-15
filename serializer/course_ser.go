package serializer

import (
	"PC/models"
)

type Course struct {
	Id           uint   `json:"id"`
	Name         string `json:"name",omitempty`
	Mark         string `json:"mark,omitempty"`
	CreatePeople string `json:"creata_people,omitempty"`
	Status       bool   `json:"status,omitempty"`
	Date         string `json:"date,omitempty"`

	SubName        string  `json:"courseName,omitempty"` //课程名称
	ClassPrice     uint    `json:"coursePrice,omitempty"`
	ClassOrPrice   uint    `json:"class_orPrice,omitempty"` //原价
	ClassTime      int64   `json:"startTime,omitempty"`
	ClassEndTime   int64   `json:"endTime,omitempty"`
	ClassStudent   uint    `json:"studentsNum,omitempty"`
	SurplusStudent uint    `json:"surplus_student"`
	Time           uint    `json:"time,omitempty"` //剩余支付时间
	BookList       []*Book `json:"book_list,omitempty"`

	SortNum uint `json:"sort_num,omitempty"` //排序

	Content string `json:"content,omitempty"` //文本内容

	View            uint64   `json:"view,omitempty"` //浏览数
	UserOrder       []*Order `json:"user_order",omitempty`
	UserOrderLength int      `json:"user_order_length"`
}

func BuildCourseSerializer(item models.Course) *Course {
	return &Course{
		Id:           item.ID,
		Name:         item.Name,
		Mark:         item.Mark,
		CreatePeople: item.CreatePeople,
		Status:       item.FormatStatus(),
		Date:         FormatTimeRFC(item.CreatedAt), //创建日期
		SubName:      item.SubName,
		ClassPrice:   item.ClassPrice,
		ClassOrPrice: item.ClassOrPrice,
		//时间戳
		ClassTime:       item.ClassTime,
		ClassEndTime:    item.ClassEndTime,
		ClassStudent:    item.ClassStudent,
		SurplusStudent:  item.SurplusStudent,
		Time:            60000,
		BookList:        FormatBookList(item.Book),
		SortNum:         item.SortNum,
		Content:         item.Content,
		View:            item.View(),
		UserOrder:       FormatSucceOrders(item.UserOrder),
		UserOrderLength: len(item.UserOrder), //一个课程包含的订单数
	}
}

func BuildCourseSerializers(items []*models.Course) (courses []*Course) {

	for _, v := range items {
		courses = append(courses, BuildCourseSerializer(*v))
	}
	return

}
