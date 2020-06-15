package adminuser

import (
	"PC/models"
	"PC/serializer"
)

type AdminUser struct {
	Page int `form:"page" json:"page" `
	Size int `form:"size" json:"size" `
}

func (item *AdminUser) GetAdminUsers() serializer.Response {

	var users []*models.User

	models.DB.Preload("Address").Find(&users)

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildAdminUsersSerializer(users),
	}

}
