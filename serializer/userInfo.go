package serializer

import "PC/models"

type UserInfo struct {
	Id        uint   `json:"id"`
	OpenId    string `json:"openId"`
	NickName  string `json:"nick_name"`
	IPAddress string `json:"ip_address"`
}

func BuildUserInfoSerializer(item models.User) *UserInfo {
	return &UserInfo{
		Id:        item.ID,
		OpenId:    item.OpenId,
		NickName:  item.Name,
		IPAddress: item.IPAddress,
	}
}
