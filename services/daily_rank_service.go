package service

import (
	"PC/cache"
	"PC/models"
	"PC/serializer"
	"fmt"
	"strings"
)

// DailyRankService 每日排行的服务
type DailyRankService struct {
}

// Get 获取排行
func (service *DailyRankService) GetCourseRank() serializer.Response {
	var courses []*models.Course

	// 从redis读取点击前十的视频
	vids, _ := cache.RedisClient.ZRevRange(cache.DailyRankKey, 0, 9).Result()

	fmt.Println(vids, "--------vids-------")
	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id, %s)", strings.Join(vids, ","))
		err := models.DB.Where("id in (?)", vids).Order(order).Find(&courses).Error
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}
	}

	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildCourseSerializers(courses),
	}
}
