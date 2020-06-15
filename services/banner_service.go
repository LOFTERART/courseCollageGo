package service

import (
	"PC/models"
	"PC/serializer"
)

//前端传递参数
type Banner struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	SortNum  uint    `json:"sort_num"`
	Link     *string `json:"link"`
	Status   uint    `json:"status"`
	LocalNum *uint   `json:"local_num"`
	Content  *string `json:"content"`
	IMG      string  `json:"img"`
}

//增加banner
func (item *Banner) CreateBanner() serializer.Response {

	info := models.Banner{
		Name:     item.Name,
		Link:     item.Link,
		Status:   item.Status,
		SortNum:  item.SortNum,
		LocalNum: item.LocalNum,
		Content:  item.Content,
		IMG:      item.IMG,
	}
	models.DB.Create(&info)

	return serializer.Response{
		Code: 20000,
		Data: nil,
		Msg:  "创建成功",
	}

}

//获取banner ALL
func (item *Banner) GetBanner() serializer.Response {

	var info []*models.Banner
	if item.Status == uint(3) {
		models.DB.Order("sort_num desc").Find(&info)
	} else {
		models.DB.Where("status=?", item.Status).Order("sort_num desc").Find(&info)
	}

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildBannerSSerializers(info),
		Msg:  "获取成功",
	}

}

//获取banner ONE
func (item *Banner) GetBannerONE() serializer.Response {

	var info models.Banner

	models.DB.First(&info, item.ID)

	info.AddView()

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildBannerSerializer(info),
		Msg:  "获取成功",
	}

}

//更新banner
func (item *Banner) UpdataBanner() serializer.Response {

	info := models.Banner{
		Name:     item.Name,
		Link:     item.Link,
		Status:   item.Status,
		SortNum:  item.SortNum,
		LocalNum: item.LocalNum,
		Content:  item.Content,
		IMG:      item.IMG,
	}

	models.DB.Where("id=?", item.ID).Model(&info).Updates(&info)

	return serializer.Response{
		Code: 20000,
		Data: nil,
		Msg:  "更新成功",
	}

}
