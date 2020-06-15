package models

import (
	"github.com/jinzhu/gorm"
	"os"
)

type Index struct {
	gorm.Model
	Name            string
	SubName         string
	ActivePic       string
	PinPic          string
	Content         string `gorm:"type:text"`
	ActiveStartTime int64  //时间戳
	ActiveEndTime   int64  //时间戳
	CreatePeople    string

	PinPrice   uint
	PinOrPrice uint
	PinDesc    string
}

func (item *Index) FormatPic(pic string) string {

	return os.Getenv("IMGADDRESS") + pic

}
