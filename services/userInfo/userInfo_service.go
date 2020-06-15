package userInfo

import (
	"PC/models"
	"PC/serializer"
	"PC/utils"
	"github.com/jinzhu/gorm"
	"github.com/medivhzhan/weapp/v2"
	"os"
)

//通过code 获取openid
type UserInfo struct {
	Code string `form:"code" json:"code" binding:"required"`
}

func (u *UserInfo) AddUser(ip string) serializer.Response {

	res, _ := weapp.Login(os.Getenv("APPID"), os.Getenv("Secret"), u.Code)

	//构建user实体
	user := models.User{
		OpenId:    res.OpenID,
		Name:      utils.GetFullName(), //随机昵称
		IPAddress: ip,
	}

	isExistence := models.DB.Where("open_id=?", res.OpenID).First(&user).RecordNotFound()

	if isExistence {
		models.DB.Create(&user)
		return serializer.Response{
			Code: 20000,
			Data: serializer.BuildUserInfoSerializer(user),
		}
	} else {
		return serializer.Response{
			Code: 20000,
			Data: serializer.BuildUserInfoSerializer(user),
		}
	}

}

//拼单完成后提交个人信息

type UserAddress struct {
	OrderID     uint   `json:"orderID"`
	UserID      uint   `json:"id"`
	TrueName    string `json:"trueName"`
	TelPhone    string `json:"telPhone"`
	FullAddress string `json:"fullAddress"`
	Province    string `json:"province"`
	City        string `json:"city"`
	County      string `json:"County"`
}

func (addr *UserAddress) UpdateAddress() serializer.Response {

	address := models.Address{
		Province:    addr.Province,
		City:        addr.City,
		County:      addr.County,
		FullAddress: addr.FullAddress,
		TrueName:    addr.TrueName,
		TelPhone:    addr.TelPhone,
		UserID:      addr.UserID,
	}

	models.DB.Create(&address)

	orderInfo := models.UserOrder{
		Model: gorm.Model{
			ID: addr.OrderID,
		},
	}
	models.DB.Model(&orderInfo).Update("address_id", address.ID)

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "ok地址成功",
		Error: "",
	}

}

//更新订阅信息

type UserSubmit struct {
	ID uint `json:"id"`
}

func (item *UserSubmit) UpdataSubmitCode() serializer.Response {

	var user models.User

	models.DB.First(&user, item.ID).Model(&user).Update("is_submit", 2)

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "订阅成功",
		Error: "",
	}

}
