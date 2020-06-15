package classInfo

import (
	"PC/models"
	"PC/serializer"
	"PC/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/medivhzhan/weapp/payment"
	"os"
	"time"
)

//提交订单信息
type ClassInfo struct {
	CourseId uint   `form:"courseId" json:"courseId" binding:"required"`
	UserId   uint   `form:"userId" json:"userId" binding:"required"`
	NickName string `form:"nickName" json:"nickName" binding:"required"`
}

func (UserOrder *ClassInfo) GetClassInfo() serializer.Response {

	//获取课程信息
	courseDetail := models.Course{
		Model: gorm.Model{
			ID: UserOrder.CourseId,
		},
	}

	models.DB.Find(&courseDetail)
	models.DB.Model(&courseDetail).Related(&courseDetail.Book, "Book")

	courseDetail.AddView()
	//查询订单是否已经存在
	var order models.UserOrder

	models.DB.Where("course_id =? AND user_id=? ", UserOrder.CourseId, UserOrder.UserId).First(&order)
	//已经下过订单直接返回订单信息 "isHave":true,
	if order.ID != 0 {

		//是否已经支付
		var isHave bool
		if order.PayOk == 2 {
			isHave = true
		} else {
			isHave = false
		}

		return serializer.Response{
			Code: 0,
			Data: gin.H{
				"courseInfo": serializer.BuildCourseSerializer(courseDetail),
				"isHave":     isHave,
				"orderId":    order.ID,
			},
		}
	} else {
		//创建订单
		UserOrder := models.UserOrder{
			UserId:           UserOrder.UserId,
			CourseId:         UserOrder.CourseId,
			Name:             UserOrder.NickName,
			Avatar:           utils.GetImg(),
			ExpressCreatedAt: time.Now(),
		}
		models.DB.Create(&UserOrder)
		//获取书籍
		return serializer.Response{
			Code: 0,
			Data: gin.H{
				"courseInfo": serializer.BuildCourseSerializer(courseDetail),
				"isHave":     false,
				"orderId":    UserOrder.ID,
			},
		}

	}

}

//生成微信订单

type WeChatOrder struct {
	OrderId  uint   `json:"orderId"` //订单号 更新支付生成的订单号
	Body     string `json:"body"`    //商品描述
	OpenID   string `json:"open_id"` // 获取openid
	TotalFee int    `json:"total_fee"`
}

func (item *WeChatOrder) CreateWeChatOrder(ip string) serializer.Response {
	//// 新建支付订单
	form := payment.Order{
		// 必填
		AppID:      os.Getenv("APPID"),
		MchID:      os.Getenv("MCHID"),
		NotifyURL:  os.Getenv("NotifyURL"),
		Body:       item.Body,
		OpenID:     item.OpenID,
		OutTradeNo: time.Now().Format("20060102150405") + utils.RandStringRunes(18), //32位订单号 日期和随机字符串
		//TotalFee:   item.TotalFee,
		TotalFee: 1,
		IP:       ip,
	}

	//更新商户订单号OutTradeNo
	order := models.UserOrder{
		Model: gorm.Model{
			ID: item.OrderId,
		},
	}

	err := models.DB.Model(&order).Where("id = ?", item.OrderId).Update("merchant_order", form.OutTradeNo).Error
	if err != nil {
		return serializer.Response{
			Code: 200001,
			Data: nil,
			Msg:  "OutTradeNo更新失败",
		}
	}

	fmt.Println(form, "----------form提交表单-------")
	res, err := form.Unify(os.Getenv("keyScrect"))
	fmt.Println(res, "----------res提交表单-------")

	if err != nil {
		// handle error
		fmt.Println(err, "----------error------")
	}
	params, _ := payment.GetParams(res.AppID, os.Getenv("keyScrect"), res.NonceStr, res.PrePayID)

	return serializer.Response{
		Code: 20000,
		Data: gin.H{
			"timeStamp": params.Timestamp,
			"nonceStr":  params.NonceStr,
			"package":   params.Package,
			"paySign":   params.PaySign,
		},
	}

}
