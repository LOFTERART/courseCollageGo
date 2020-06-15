package orderList

import (
	"PC/cache"
	"PC/models"
	"PC/serializer"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/medivhzhan/weapp/v2"
	"os"
	"time"
)

type OrderListService struct {
	SendType uint   `json:"send_type" form:"send_type"  `
	OrderId  uint   `json:"order_id" form:"order_id"  `
	ExName   string `json:"exName" form:"exName"`     //快递名字
	ExRegion string `json:"exRegion" form:"exRegion"` //快递单号
	Operator string `json:"operator" form:"operator"` //操作人
}

func (UserOrder *OrderListService) GetOrderList() serializer.Response {
	var list []*models.UserOrder
	var course []*models.Course
	models.DB.Raw("SELECT id,name,time,avatar FROM user_order WHERE pay_ok <= 1 AND id >= ((SELECT MAX(id) FROM user_order)-(SELECT MIN(id) FROM user_order)) * RAND() + (SELECT MIN(id) FROM user_order)  LIMIT 2").
		Scan(&list)

	models.DB.Select("id,name,surplus_student").Order("sort_num desc").Find(&course)

	var homeInfo []*models.Index
	models.DB.Find(&homeInfo)

	//订单数
	var numOrderAll []models.UserOrder

	models.DB.Select("id").Find(&numOrderAll)

	return serializer.Response{
		Code: 0,
		Data: map[string]interface{}{
			"pingList": serializer.BuildClassSerializers(list),
			"actions":  serializer.BuildCourseSerializers(course),
			"homeInfo": serializer.BuildIndexSerializers(homeInfo),
			"PingNum":  len(numOrderAll),
		},
		Msg:   "SUCCESS",
		Error: "",
	}

}

//更新发货状态
func (UserOrder *OrderListService) UpOrderSend() serializer.Response {

	token, _ := cache.RedisClient.Get(cache.WeChatAccessToken).Result()

	if len(token) <= 0 {
		res, err := weapp.GetAccessToken(os.Getenv("APPID"), os.Getenv("Secret"))
		if err != nil {
			// 处理一般错误信息
			return serializer.Response{
				Code: 20000,
				Msg:  "获取微信token失败",
			}
		}

		if err := res.GetResponseError(); err != nil {
			// 处理微信返回错误信息
			return serializer.Response{
				Code: 20000,
				Msg:  "获取微信token失败",
			}
		}

		fmt.Printf("返回结果: %#v", res.AccessToken)

		// 存储微信接口凭证到redis
		cache.RedisClient.Set(cache.WeChatAccessToken, res.AccessToken, 1*time.Hour)
		token, _ = cache.RedisClient.Get(cache.WeChatAccessToken).Result()
	}

	order := models.UserOrder{
		Model: gorm.Model{
			ID: UserOrder.OrderId,
		},
	}

	err := models.DB.Model(&order).Updates(models.UserOrder{
		Express:          UserOrder.ExName,
		ExpressNumber:    UserOrder.ExRegion,
		SendType:         UserOrder.SendType,
		ExpressCreatedAt: time.Now(), //快递发单日期
		Operator:         UserOrder.Operator,
	}).First(&order).Error

	if err != nil {
		fmt.Println(err, "err")
		return serializer.Response{
			Code:  200,
			Data:  nil,
			Msg:   "更新失败",
			Error: "",
		}
	}

	//获取用户信息
	var user models.User
	models.DB.First(&user, order.UserId)

	//发货成功 发送通知
	sender := weapp.SubscribeMessage{
		ToUser:           user.OpenId,
		TemplateID:       os.Getenv("TEMPLATE"),
		Page:             "page/index/index",
		MiniprogramState: weapp.MiniprogramStateDeveloper, // 或者: "developer"
		Data: weapp.SubscribeMessageData{
			"time1": {
				Value: time.Now().Format("2006-01-02 15:04:05"),
			},
			"thing2": {
				Value: "2020年美术世界拼团课程赠送图书资料",
			},
			"name3": {
				Value: UserOrder.ExName,
			},
			"number4": {
				Value: UserOrder.ExRegion,
			},
		},
	}

	res, err := sender.Send(token)
	if err != nil {
		// 处理一般错误信息
		return serializer.Response{
			Code:  20000,
			Msg:   "微信通知消息发送失败",
			Error: err.Error(),
		}
	}

	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		return serializer.Response{
			Code:  20000,
			Msg:   "微信通知消息发送失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Code: 20000,
		Data: gin.H{
			"type": 2,
		},
		Msg:   "更新成功",
		Error: "",
	}
}
