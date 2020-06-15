package models

import (
	"github.com/jinzhu/gorm"
	"os"
)

type Book struct {
	gorm.Model
	Title        string
	Tip          string
	Img          string
	CreatePeople string //创建人
	Course       Course `gorm:"ForeignKey:CourseID"`
	CourseID     uint
}

func (item *Book) FormatPic() string {

	return os.Getenv("IMGADDRESS") + item.Img

}
