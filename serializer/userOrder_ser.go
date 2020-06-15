package serializer

import (
	"PC/models"
)

type Order struct {
	ID uint `json:"id"`

	UserId uint `json:"user_id"`

	Course *Course `json:"course"`

	CourseId uint `json:"course_id"`

	PayOk uint `json:"pay_ok"` //支付状态1 未支付 2 成功支付

	Avatar string `json:"avatar"`
	Name   string `json:"name"`
	Time   int    `json:"time"`

	Address *Address `json:"address"`

	AddressID uint `json:"address_id"`

	//	是否发货
	SendType uint `json:"send_type"` //发货状态1 未发货 2 发货

	Express string `json:"express"` //快递

	ExpressNumber string `json:"express_number"` //快递单号

	//
	ExpressCreatedAt string `json:"express_created_at"`

	Operator string `json:"operator"` //操作人

	MerchantOrder string `json:"merchant_order"` //商户订单号

	CreatedAt string `json:"created_at"`


}

func FormatSucceOrder(item *models.UserOrder) *Order {

	return &Order{
		ID:               item.ID,
		UserId:           item.UserId,
		Course:           BuildCourseSerializer(item.Course),
		CourseId:         item.CourseId,
		PayOk:            item.PayOk,
		Avatar:           item.Avatar,
		Name:             item.Name,
		Time:             item.Time,
		Address:          FormatAddress(item.Address),
		AddressID:        item.AddressID,
		SendType:          item.SendType,
		Express:          item.Express,
		ExpressNumber:    item.ExpressNumber,
		ExpressCreatedAt: item.ExpressCreatedAt.Format("2006-01-02 15:04:05"),
		Operator:         item.Operator,
		MerchantOrder:    item.MerchantOrder,
		CreatedAt:item.CreatedAt.Format("2006-01-02 15:04:05"),
	}

}

func FormatSucceOrders(item []*models.UserOrder) (items []*Order) {

	for _, v := range item {
		items = append(items, FormatSucceOrder(v))
	}
	return

}
