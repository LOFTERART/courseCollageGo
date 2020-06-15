package getSuccessOrder

import (
	"PC/models"
	"PC/serializer"
)

type SuccessOrder struct {
	Page int `form:"page" json:"page" `
	Size int `form:"size" json:"size" `
	//查询是否已经发货状态
	SendType uint   `json:"send_type"`
	TelPhone string `json:"tel_phone"`
}

func (order *SuccessOrder) GetSuccessOrder() serializer.Response {

	var orderList []*models.UserOrder

	total := 0

	if order.Size == 0 {
		order.Size = 5
	}
	if order.Page == 0 {
		order.Page = 1
	}

	start := (order.Page - 1) * order.Size

	var addr []models.Address
	if len(order.TelPhone) > 0 {

		models.DB.Where("tel_phone LIKE ? ", "%"+order.TelPhone+"%").Find(&addr)
		var address []uint
		for _, v := range addr {
			address = append(address, v.ID)
		}

		models.DB.Where(" address_id IN (?) AND pay_ok=?", address, 2).Find(&orderList).Count(&total)
		models.DB.Where(" address_id IN (?) AND pay_ok=?", address, 2).Preload("Address").Preload("Course").Find(&orderList)
	} else {
		models.DB.Where("pay_ok=? AND send_type =?", 2, order.SendType).Find(&orderList).Count(&total)
		models.DB.Where("pay_ok=? AND send_type =?", 2, order.SendType).Preload("Address").Preload("Course").Limit(order.Size).Offset(start).Find(&orderList)
	}

	return serializer.BuildListResponse(serializer.FormatSucceOrders(orderList), uint(total), 20000)

}
