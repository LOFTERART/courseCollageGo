package service

import (
	"PC/models"
	"PC/serializer"
)

type WeChatIndex struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	SubName         string `json:"sub_name"`
	ActivePic       string `json:"active_pic"`
	PinPic          string `json:"pin_pic"`
	Content         string `json:"content"`
	ActiveStartTime int64  `json:"active_start_time"` //时间戳
	ActiveEndTime   int64  `json:"active_end_time"`   //时间戳
	CreatePeople    string `json:"create_people"`

	PinPrice   uint   `json:"pin_price"`
	PinOrPrice uint   `json:"pin_or_price"`
	PinDesc    string `json:"pin_desc"`
}

func (item *WeChatIndex) CreateWeChatIndex() serializer.Response {

	info := models.Index{
		Name:            item.Name,
		SubName:         item.SubName,
		ActivePic:       item.ActivePic,
		PinPic:          item.PinPic,
		Content:         item.Content,
		ActiveStartTime: item.ActiveStartTime,
		ActiveEndTime:   item.ActiveEndTime,
		CreatePeople:    item.CreatePeople,
		PinPrice:        item.PinPrice,
		PinOrPrice:      item.PinOrPrice,
		PinDesc:         item.PinDesc,
	}
	if item.ID > 0 {
		models.DB.Where("id=?", item.ID).Model(&info).Updates(&info)
		return serializer.Response{
			Code:  20000,
			Data:  nil,
			Msg:   "更新成功",
			Error: "",
		}

	} else {
		models.DB.Create(&info)
		return serializer.Response{
			Code:  20000,
			Data:  nil,
			Msg:   "创建成功",
			Error: "",
		}
	}

}

//获取首页活动信息
func (item *WeChatIndex) GetWeChatIndex() serializer.Response {

	var info []*models.Index
	models.DB.Limit(1).Find(&info)
	return serializer.Response{
		Code:  20000,
		Data:  serializer.BuildIndexSerializers(info),
	}

}
