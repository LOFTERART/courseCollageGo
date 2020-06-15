package serializer

import "PC/models"

type Users struct {
	Id      uint       `json:"id"`
	Name    string     `json:"name"`
	Address []*Address `json:"address"`
}

//格式化单个用户
func BuildAdminUserSerializer(item *models.User) *Users {
	return &Users{
		Id:      item.ID,
		Name:    item.Name,
		Address: FormatAddressList(item.Address),
	}
}

//格式化多个用户
func BuildAdminUsersSerializer(item []*models.User) (items []*Users) {

	for _, v := range item {
		items = append(items, BuildAdminUserSerializer(v))
	}
	return

}
