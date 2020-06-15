package tasks

import (
	"PC/cache"
	"PC/models"
	"fmt"
	"github.com/medivhzhan/weapp/v2"
	"os"
	"time"
)

// RestartDailyRank 重启一天的排名
func RestartDailyRank() error {
	return cache.RedisClient.Del("rank:daily").Err()
}

//是否发送消息  判断活动时间戳
func SendMessage() error {

	fmt.Println("-------开始执行了------定时任务")
	//获取当前时间戳
	currentTime := time.Now().UnixNano() / 1e6
	//获取活动开始时间

	var info models.Index

	models.DB.First(&info)
	fmt.Println(info.ActiveEndTime)

	//如果当前时间大于活动开始时间 发送订阅通知
	if currentTime > info.ActiveStartTime {

		Sendmsg()
	}

	return nil
}

//发送通知
func Sendmsg() {
	//获取token
	token, _ := cache.RedisClient.Get(cache.WeChatAccessToken).Result()
	if len(token) <= 0 {
		res, err := weapp.GetAccessToken(os.Getenv("APPID"), os.Getenv("Secret"))
		if err != nil {
			// 处理一般错误信息
			fmt.Println(err, "--------发送通知-----")
		}

		if err := res.GetResponseError(); err != nil {
			// 处理微信返回错误信息
			fmt.Println(err, "--------发送通知-----")
		}

		// 存储微信接口凭证到redis
		cache.RedisClient.Set(cache.WeChatAccessToken, res.AccessToken, 1*time.Hour)
		token, _ = cache.RedisClient.Get(cache.WeChatAccessToken).Result()
	}

	//	获取用户
	var Users []*models.User
	models.DB.Where("is_submit=?", 2).Find(&Users)

	var u models.User
	for _, v := range Users {

		//发货成功 发送通知
		sender := weapp.SubscribeMessage{
			ToUser:           v.OpenId,
			TemplateID:       os.Getenv("TEMPLATESUBMIT"),
			Page:             "page/index/index",
			MiniprogramState: weapp.MiniprogramStateDeveloper, // 或者: "developer"
			Data: weapp.SubscribeMessageData{
				"thing1": {
					Value: "XX机构拼团",
				},
				"thing2": {
					Value: "订阅的拼团项目活动已经开始",
				},
				"thing7": {
					Value: "学校班车免费接送参观校区",
				},
			},
		}

		res, err := sender.Send(token)
		if err != nil {
			// 处理一般错误信息
			fmt.Println(err, "订阅通知1")

		}

		if err := res.GetResponseError(); err != nil {
			// 处理微信返回错误信息
			fmt.Println(err, "订阅通知2")
			return
		}

		//	修改用户状态

		models.DB.First(&u, v.ID).Model(&u).Update("is_submit", 1)

	}

}
