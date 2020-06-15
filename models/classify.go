package models

import (
	"github.com/jinzhu/gorm"
	"os"
)

type Classify struct {
	gorm.Model
	Name    string
	SortNum uint //排序
	IMG     string
	Article []*Article
}

func (item *Classify) FormatPic() string {

	return os.Getenv("IMGADDRESS") + item.IMG

}
