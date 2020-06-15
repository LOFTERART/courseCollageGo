package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	//省市县
	Province string
	City     string
	County   string
	//详细地址
	FullAddress string
	//	真实姓名
	TrueName string
	TelPhone string

	//用户id
	UserID uint
}
