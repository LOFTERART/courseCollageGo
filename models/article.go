package models

import (
	"PC/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Article struct {
	gorm.Model
	Name    string
	SubName string
	Content string `gorm:"type:text"`
	Views   uint
	//	关联classify
	ClassifyID uint
	Classify   Classify
}

// item 点击数
func (item *Article) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.ArticleViewKey(item.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//增加课程浏览量

func (item *Article) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.ArticleViewKey(item.ID))

}
