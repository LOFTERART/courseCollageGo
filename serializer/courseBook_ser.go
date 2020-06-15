package serializer

import "PC/models"

type Book struct {
	Id           uint    `json:"id"`
	Title        string  `json:"title"`
	Tip          string  `json:"tip"`
	Img          string  `json:"img"`
	CreatePeople string  `json:"create_people"`
	CourseID     uint    `json:"course_id"`
	Course       *Course `json:"course"`
	ImgURL string `json:"img_url"`
}

func FormatBook(item models.Book) *Book {
	return &Book{
		Id:           item.ID,
		Title:        item.Title,
		Tip:          item.Tip,
		Img:          item.Img,
		ImgURL:          item.FormatPic(),
		CreatePeople: item.CreatePeople,
		CourseID:     item.CourseID,
		Course:       BuildCourseSerializer(item.Course),
	}
}

func FormatBookList(item []models.Book) (books []*Book) {
	for _, v := range item {
		books = append(books, FormatBook(v))
	}

	return

}
