package address

import (
	"PC/models"
	"PC/serializer"
	"github.com/jinzhu/gorm"
)

//获取地址列表 用户ID
type Address struct {
	UserId uint `json:"userId"`
}

func (addr *Address) GetAddressLists() serializer.Response {
	var addrs []*models.Address
	models.DB.Where("user_id=?", addr.UserId).Find(&addrs)
	return serializer.Response{
		Code:  0,
		Data:  serializer.FormatAddressList(addrs),
		Msg:   "",
		Error: "",
	}

}

//用户提交地址id 更新订单的address_id信息

type UserChoseAddr struct {
	AddressId uint `json:"AddressId"`
	OderId    uint `json:"oder_id"`
}

func (uAddr *UserChoseAddr) UpdataOrderAddress() serializer.Response {

	order := models.UserOrder{
		Model: gorm.Model{
			ID: uAddr.OderId,
		},
	}

	models.DB.Model(&order).Update("address_id", uAddr.AddressId)

	return serializer.Response{
		Data: "提交订单ok",
	}

}
