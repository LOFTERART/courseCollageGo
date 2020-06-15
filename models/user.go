package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name      string `gorm:"size:20"`
	OpenId    string `gorm:"size:100"`
	IPAddress string `gorm:"size:20"`
	ISSubmit  uint   `gorm:"default: 1 "` //是否订阅活动 发送通知 未订阅 1   订阅2

	UserOrder []*UserOrder `gorm:"ForeignKey:UserId"`
	Address   []*Address   `gorm:"ForeignKey:UserId"`
}
