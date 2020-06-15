package api

import (
	"PC/utils"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
)

func UploadImg(c *gin.Context) {

	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	var name string
	for _, file := range files {
		ext := strings.ToLower(path.Ext(file.Filename))
		name = utils.RandStringRunes(10) + ext
		// 上传文件至指定目录
		c.SaveUploadedFile(file, "./uploadImg/"+name)

	}

	c.JSON(200, gin.H{
		"code": 20000,
		"msg":  "ok",
		"files": gin.H{
			"file":    os.Getenv("IMGADDRESS") + name,
			"imgName": name,
		},
	})

}
