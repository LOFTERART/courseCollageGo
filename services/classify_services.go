package service

import (
	"PC/models"
	"PC/serializer"
)

//前端传递参数
type Classify struct {
	Name    string `json:"name"`
	SortNum uint   `json:"sort_num"`
	ID      uint   `json:"id"`
	IMG string `json:"img"`
}

//增加分类
func (item *Classify) CreateClassify() serializer.Response {

	info := models.Classify{Name: item.Name, SortNum: item.SortNum,IMG: item.IMG}
	models.DB.Create(&info)

	return serializer.Response{
		Code: 20000,
		Data: nil,
		Msg:  "创建成功",
	}

}

//获取分类
func (item *Classify) GetClassify() serializer.Response {

	var info []*models.Classify

	models.DB.Order("sort_num desc").Find(&info)

	return serializer.Response{
		Code: 20000,
		Data: serializer.FormatClassifyList(info),
		Msg:  "创建成功",
	}

}

//更新分类
func (item *Classify) UpdataClassifyS() serializer.Response {

	info := models.Classify{Name: item.Name, SortNum: item.SortNum,IMG: item.IMG}

	models.DB.Where("id=?", item.ID).Model(&info).Updates(&info)

	return serializer.Response{
		Code: 20000,
		Data: nil,
		Msg:  "更新成功",
	}

}

//根据分类获取文章
type ClassifyArticle struct {
	ClassifyID uint `json:"classify_id"`
}

func (item *ClassifyArticle) GetClassifyArticle() serializer.Response {

	var articles []*models.Article

	models.DB.Where("classify_id=?", item.ClassifyID).Find(&articles)

	return serializer.Response{
		Code: 20000,
		Data: serializer.FormatArticleList(articles),
	}

}
