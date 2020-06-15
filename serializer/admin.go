package serializer

import (
	"PC/models"
)

type Admin struct {
	Id     uint   `json:"id"`
	Name   string `json:"name"`
	Tokens string `json:"token"`
}

func BuildAdminSerializer(item models.Admin) *Admin {
	return &Admin{
		Id:     item.ID,
		Name:   item.Name,
		Tokens: item.Tokens,
	}
}

func BuildAdminSerializers(item []*models.Admin) (items []*Admin) {

	for _, v := range item {
		items = append(items, BuildAdminSerializer(*v))
	}
	return

}
