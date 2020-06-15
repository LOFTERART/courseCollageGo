package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserOrder struct {
	gorm.Model

	UserId uint

	Course Course `gorm:"ForeignKey:CourseId"`

	CourseId uint

	PayOk uint `gorm:"default: 1 "` //支付状态1 未支付 2 成功支付

	Avatar string
	Name   string
	Time   int `gorm:"default: 60000 "`

	Address Address `gorm:"ForeignKey:AddressID"`

	AddressID uint

	//	是否发货
	SendType uint `gorm:"default: 1 "` //发货状态1 未发货 2 发货

	Express string //快递

	ExpressNumber string //快递单号

	//
	ExpressCreatedAt time.Time

	Operator string `gorm:"default: '未知' "` //操作人

	MerchantOrder string //商户订单号
}
