package models

import (
	"PC/cache"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
)

type Banner struct {
	gorm.Model
	Name     string
	Link     *string
	Status   uint    `gorm:"default: 1 "` //是否显示 1显示 2隐藏
	SortNum  uint    //排序 数字越大越靠前
	LocalNum *uint   //本地文章id
	Content  *string `gorm:"type:text"`
	IMG      string
}

func (item *Banner) FormatPic(pic string) string {

	return os.Getenv("IMGADDRESS") + pic

}

// item 点击数
func (item *Banner) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.BannerViewKey(item.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//增加课程浏览量

func (item *Banner) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.BannerViewKey(item.ID))

}
