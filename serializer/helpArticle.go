package serializer

import "PC/models"

type Article struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	SubName    string    `json:"sub_name"`
	Content    string    `json:"content"`
	Views      uint64    `json:"views"`
	CreateTime string    `json:"create_time"`
	Classify   *Classify `json:"classify"`
}

func FormatArticle(item *models.Article) *Article {
	return &Article{
		Id:         item.ID,
		Name:       item.Name,
		SubName:    item.SubName,
		Content:    item.Content,
		Classify:   FormatClassify(item.Classify),
		Views:      item.View(),
		CreateTime: item.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func FormatArticleList(item []*models.Article) (books []*Article) {
	for _, v := range item {
		books = append(books, FormatArticle(v))
	}

	return

}
