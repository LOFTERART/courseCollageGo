package service

import (
	"PC/models"
	"PC/serializer"
)

//前端传递参数
type Studio struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Tel       string  `json:"tel"`
	SortNum   uint    `json:"sort_num"` //排序 数字越大越靠前
	Address   string  `json:"address"`  //本地文章id
	Content   string  `json:"content"`
	IMG       string  `json:"img"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Status    uint    `json:"status"`
}

//增加Studio
func (item *Studio) CreateStudio() serializer.Response {

	info := models.Studio{
		Name:      item.Name,
		Tel:       item.Tel,
		SortNum:   item.SortNum,
		Address:   item.Address,
		Content:   item.Content,
		IMG:       item.IMG,
		Longitude: item.Longitude,
		Latitude:  item.Latitude,
		Status:    item.Status,
	}
	models.DB.Create(&info)

	return serializer.Response{
		Code: 20000,
		Data: nil,
		Msg:  "创建成功",
	}

}

//获取Studio ALL
func (item *Studio) GetStudio() serializer.Response {

	var info []*models.Studio

	models.DB.Order("sort_num desc").Find(&info)

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildStudioSSerializers(info),
		Msg:  "获取成功",
	}

}

//获取Studio ONE
func (item *Studio) GetStudioOne() serializer.Response {

	var info models.Studio

	models.DB.First(&info, item.ID)

	info.AddView()

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildStudioSerializer(&info),
		Msg:  "获取成功",
	}

}

//更新Studio
func (item *Studio) UpdataStudio() serializer.Response {

	info := models.Studio{
		Name:      item.Name,
		Tel:       item.Tel,
		SortNum:   item.SortNum,
		Address:   item.Address,
		Content:   item.Content,
		IMG:       item.IMG,
		Longitude: item.Longitude,
		Latitude:  item.Latitude,
		Status:    item.Status,
	}

	models.DB.Where("id=?", item.ID).Model(&info).Updates(&info)

	return serializer.Response{
		Code: 20000,
		Data: nil,
		Msg:  "更新成功",
	}

}
