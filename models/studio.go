package models

import (
	"PC/cache"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

type Studio struct {
	gorm.Model
	Name      string
	Tel       string
	Status    uint   `gorm:"default: 1 "` //是否显示 1显示 2隐藏
	SortNum   uint   //排序 数字越大越靠前
	Address   string //本地文章id
	Content   string `gorm:"type:text"`
	IMG       string
	Longitude float64
	Latitude  float64
}

func (item *Studio) FormatPic(pic string) string {

	return os.Getenv("IMGADDRESS") + pic

}

// item 点击数
func (item *Studio) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.StudioViewKey(item.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//增加课程浏览量

func (item *Studio) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.StudioViewKey(item.ID))

}
