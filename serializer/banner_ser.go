package serializer

import (
	"PC/models"
)

type Banner struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Link     *string `json:"link"`
	Status   uint    `json:"status"`   //是否显示 1显示 2隐藏
	SortNum  uint    `json:"sort_num"` //排序 数字越大越靠前
	LocalNum *uint   `json:"local_num"`
	Content  *string `json:"content"`
	IMG      string  `json:"img"`    //图片地址数据库
	IMGURL   string  `json:"imgurl"` //显示地址http
	Views    uint64  `json:"views"`
}

func BuildBannerSerializer(item models.Banner) *Banner {
	return &Banner{
		ID:       item.ID,
		Name:     item.Name,
		Link:     item.Link,
		Status:   item.Status,
		SortNum:  item.SortNum,
		LocalNum: item.LocalNum,
		Content:  item.Content,
		IMG:      item.IMG,
		IMGURL:   item.FormatPic(item.IMG),
		Views:    item.View(),
	}

}

func BuildBannerSSerializers(item []*models.Banner) (ads []*Banner) {

	for _, v := range item {
		ads = append(ads, BuildBannerSerializer(*v))
	}
	return

}
