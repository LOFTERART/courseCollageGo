package serializer

import "PC/models"

type Studio struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Tel       string  `json:"tel"`
	SortNum   uint    `json:"sort_num"` //排序 数字越大越靠前
	Address   string  `json:"address"`  //本地文章id
	Content   string  `json:"content"`
	IMG       string  `json:"img"`
	IMGURL    string  `json:"imgurl"` //显示地址http
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Status    uint    `json:"status"`
	Views     uint64  `json:"views"`
}

func BuildStudioSerializer(item *models.Studio) *Studio {
	return &Studio{
		ID:        item.ID,
		Name:      item.Name,
		Tel:       item.Tel,
		Address:   item.Address,
		SortNum:   item.SortNum,
		Content:   item.Content,
		IMG:       item.IMG,
		Longitude: item.Longitude,
		Latitude:  item.Latitude,
		Status:    item.Status,
		Views:     item.View(),
		IMGURL:    item.FormatPic(item.IMG),
	}

}

func BuildStudioSSerializers(item []*models.Studio) (items []*Studio) {

	for _, v := range item {
		items = append(items, BuildStudioSerializer(v))
	}
	return

}
