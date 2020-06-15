package serializer

import (
	"PC/models"
)

type UserOrder struct {
	Id       uint   `json:"id"`
	Avatar   string `json:"avatar,omitempty"`
	Name     string `json:"name,omitempty"`
	Time     int    `json:"time,omitempty"`
	CourseId uint   `json:"course_id,omitempty"`
	PayOk    uint   `json:"pay_ok,omitempty"`

	UserId uint `json:"user_id,omitempty"`
}

func BuildClassSerializer(item models.UserOrder) *UserOrder {
	return &UserOrder{
		Id:       item.ID,
		Avatar:   item.Avatar,
		Name:     item.Name,
		Time:     item.Time,
		CourseId: item.CourseId,
		PayOk:    item.PayOk,
		UserId:   item.UserId,
	}
}

func BuildClassSerializers(item []*models.UserOrder) (ads []*UserOrder) {

	for _, v := range item {
		ads = append(ads, BuildClassSerializer(*v))
	}
	return

}
