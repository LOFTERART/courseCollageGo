package serializer

import "PC/models"

type Index struct {
	Id              uint   `json:"id"`
	Name            string `json:"name"`
	SubName         string `json:"sub_name"`
	ActivePic       string `json:"active_pic"`
	PinPic          string `json:"pin_pic"`
	ActivePicStr    string `json:"active_pic_str"`
	PinPicStr       string `json:"pin_pic_str"`
	Content         string `gorm:"type:text" json:"content"`
	ActiveStartTime int64  `json:"active_start_time"` //时间戳
	ActiveEndTime   int64  `json:"active_end_time"`   //时间戳
	CreatePeople    string `json:"create_people"`

	PinPrice   uint   `json:"pin_price"`
	PinOrPrice uint   `json:"pin_or_price"`
	PinDesc    string `json:"pin_desc"`
}

func BuildIndexSerializer(item models.Index) *Index {
	return &Index{
		Id:              item.ID,
		Name:            item.Name,
		SubName:         item.SubName,
		ActivePic:       item.FormatPic(item.ActivePic),
		PinPic:          item.FormatPic(item.PinPic),
		ActivePicStr:    item.ActivePic,
		PinPicStr:       item.PinPic,
		Content:         item.Content,
		ActiveStartTime: item.ActiveStartTime,
		ActiveEndTime:   item.ActiveEndTime,
		CreatePeople:    item.CreatePeople,

		PinPrice:   item.PinPrice,
		PinOrPrice: item.PinOrPrice,
		PinDesc:    item.PinDesc,
	}
}

func BuildIndexSerializers(item []*models.Index) (items []*Index) {

	for _, v := range item {
		items = append(items, BuildIndexSerializer(*v))
	}
	return

}
