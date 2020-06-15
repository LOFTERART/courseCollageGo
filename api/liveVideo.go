package api

import (
	"PC/cache"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/medivhzhan/weapp/v2"
	"log"
	"os"
	"time"
)

const (
	BASEURL     = "https://api.weixin.qq.com/wxaapi/broadcast/room/create?access_token="
	WXAPPID     = "wxdb7814ea11e729c0"
	WXAppSecret = "6e2ffebd2720bcc86dfeb6bb1a54f358"
)

type Room struct {
	RoomId  uint `json:"roomId"`
	ErrCode uint `json:"errcode"`
}

//创建直播间
func GetRoomID(c *gin.Context) {

	token, _ := cache.RedisClient.Get(cache.WeChatAccessToken).Result()

	if len(token) <= 0 {
		res, err := weapp.GetAccessToken(os.Getenv("APPID"), os.Getenv("Secret"))
		if err != nil {
			// 处理一般错误信息

		}

		if err := res.GetResponseError(); err != nil {
			// 处理微信返回错误信息

		}

		// 存储微信接口凭证到redis
		cache.RedisClient.Set(cache.WeChatAccessToken, res.AccessToken, 1*time.Hour)
		token, _ = cache.RedisClient.Get(cache.WeChatAccessToken).Result()
	}

	header := req.Header{
		"Accept": "application/json",
	}

	param := req.Param{
		"name":         "测试直播房间1",                                                          // 房间名字
		"coverImg":     "zIqacGdqB_CA4knWoEjLbQn9PxxIkCX_GeQXL-mdfYOXb4nHZDTGag2EdkfEVDN_", // 通过 uploadfile 上传，填写 mediaID
		"startTime":    1591793302,                                                         // 开始时间
		"endTime":      1591796902,                                                         // 结束时间
		"anchorName":   "美术世界",                                                             // 主播昵称
		"anchorWechat": "gogo12345y",                                                       // 主播微信号
		"shareImg":     "zIqacGdqB_CA4knWoEjLbQn9PxxIkCX_GeQXL-mdfYOXb4nHZDTGag2EdkfEVDN_", //通过 uploadfile 上传，填写 mediaID
		"type":         1,                                                                  // 直播类型，1 推流 0 手机直播
		"screenType":   0,                                                                  // 1：横屏 0：竖屏
		"closeLike":    0,                                                                  // 是否 关闭点赞 1 关闭
		"closeGoods":   0,                                                                  // 是否 关闭商品货架，1：关闭
		"closeComment": 0,                                                                  // 是否开启评论，1：关闭
	}

	res, err := req.Post(BASEURL+token, header, param)
	if err != nil {
		log.Fatal(err)
	}
	var room Room
	res.ToJSON(&room) // 响应体转成对象
	fmt.Println(room.RoomId, room.ErrCode, "-----------")

}

//获取直播间

type LiveVIdeo struct {
	Start uint
	limit uint
}

func GetRoomList(c *gin.Context) {
	token, _ := cache.RedisClient.Get(cache.WeChatAccessToken).Result()

	if len(token) <= 0 {
		res, err := weapp.GetAccessToken(os.Getenv("APPID"), os.Getenv("Secret"))
		if err != nil {
			// 处理一般错误信息

		}

		if err := res.GetResponseError(); err != nil {
			// 处理微信返回错误信息

		}

		// 存储微信接口凭证到redis
		cache.RedisClient.Set(cache.WeChatAccessToken, res.AccessToken, 1*time.Hour)
		token, _ = cache.RedisClient.Get(cache.WeChatAccessToken).Result()
	}

	info := LiveVIdeo{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {

		header := req.Header{
			"Accept": "application/json",
		}
		param := req.Param{
			"start": info.Start,
			"limit": info.limit,
		}

		res, err := req.Post("https://api.weixin.qq.com/wxa/business/getliveinfo?access_token="+token, header, req.BodyJSON(&param))
		if err != nil {
			log.Fatal(err)
		}
		var room map[string]interface{}
		res.ToJSON(&room) // 响应体转成对象

		c.JSON(200, &room)

	}

}
