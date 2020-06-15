package models

import (
	"PC/cache"
	"github.com/jinzhu/gorm"
	"strconv"
)

type Course struct {
	gorm.Model
	Name         string //分类名称
	Mark         string //备注
	CreatePeople string //创建人
	Status       uint   `gorm:"default: 1 "` //启用状态 1 启用  2 未启用

	SortNum uint // 课程排序

	SubName        string //课程名称
	ClassPrice     uint   //优惠价
	ClassOrPrice   uint   //原价
	ClassTime      int64
	ClassEndTime   int64
	ClassStudent   uint   //班级名额
	SurplusStudent uint   //剩余学生人数
	Book           []Book `gorm:"ForeignKey:CourseID"`

	Content string `gorm:"type:text"`

	Views     uint64
	UserOrder []*UserOrder //添加用户订单 1对多关系
}

func (item *Course) FormatStatus() bool {
	if item.Status == 1 {
		return true
	} else {
		return false
	}
}

// item 点击数
func (item *Course) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(item.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

//增加课程浏览量

func (item *Course) AddView() {
	// 增加视频点击数
	cache.RedisClient.Incr(cache.VideoViewKey(item.ID))
	// 增加排行点击数
	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(item.ID)))
}
