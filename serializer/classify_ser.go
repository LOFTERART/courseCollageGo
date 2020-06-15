package serializer

import "PC/models"

type Classify struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	SortNum uint   `json:"sort_num"` //排序
	IMG     string `json:"img"`
	IMGURL  string `json:"imgurl"`
}

func FormatClassify(item models.Classify) *Classify {
	return &Classify{
		ID:      item.ID,
		Name:    item.Name,
		SortNum: item.SortNum,
		IMG:     item.IMG,
		IMGURL:  item.FormatPic(),
	}
}

func FormatClassifyList(item []*models.Classify) (books []*Classify) {
	for _, v := range item {
		books = append(books, FormatClassify(*v))
	}

	return

}
