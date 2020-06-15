package service

import (
	"PC/models"
	"PC/serializer"
	"fmt"
)

type Article struct {
	Id         uint   `form:"id" json:"id"` //文章id
	Name       string `form:"name" json:"name"`
	SubName    string `form:"sub_name" json:"sub_name"`
	Content    string `form:"content" json:"content"`
	ClassifyID uint   `form:"classify_id" json:"classify_id"` //分类id
	IsEdit     bool   `form:"is_edit" json:"is_edit"`
}

//增加文字和编辑文章 IsEdit是否传递这个标示

func (item *Article) AddCreateArticle() serializer.Response {

	info := models.Article{
		Name:       item.Name,
		SubName:    item.SubName,
		Content:    item.Content,
		ClassifyID: item.ClassifyID,
	}

	//是编辑过来的 更新文章
	if item.IsEdit {
		fmt.Println("------AAAAA---------")
		models.DB.Model(&info).Updates(&info)
	} else {
		models.DB.Create(&info)
	}

	return serializer.Response{
		Code:  20000,
		Data:  nil,
		Msg:   "成功",
		Error: "",
	}

}

//查询多个
func (item *Article) GetCreateArticle() serializer.Response {

	var info []*models.Article

	models.DB.Preload("Classify").Find(&info)

	return serializer.Response{
		Code:  20000,
		Data:  serializer.FormatArticleList(info),
		Msg:   "查询成功",
		Error: "",
	}

}

//查询一个
func (item *Article) GetCreateArticleOne() serializer.Response {

	var info models.Article

	models.DB.First(&info, item.Id)
	info.AddView()
	return serializer.Response{
		Code:  20000,
		Data:  serializer.FormatArticle(&info),
		Msg:   "获取成功",
		Error: "",
	}

}
