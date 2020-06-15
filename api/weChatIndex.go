package api

import (
	service "PC/services"
	"PC/utils"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strings"
)

//创建首页信息
func WeChatIndex(c *gin.Context) {

	var WeChatIndex service.WeChatIndex

	if err := c.ShouldBind(&WeChatIndex); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := WeChatIndex.CreateWeChatIndex()
		c.JSON(0, &res)
	}

}

//获取首页信息 和详情信息
func WeChatIndexList(c *gin.Context) {

	var WeChatIndex service.WeChatIndex

	if err := c.ShouldBind(&WeChatIndex); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := WeChatIndex.GetWeChatIndex()
		c.JSON(0, &res)
	}

}

//图片上传
func WeChatIndexUploadImg(c *gin.Context) {

	file, _ := c.FormFile("file")

	ext := strings.ToLower(path.Ext(file.Filename))
	name := utils.RandStringRunes(10) + ext
	// 上传文件至指定目录
	c.SaveUploadedFile(file, "./uploadImg/"+name)

	c.JSON(200, gin.H{
		"code": 20000,
		"msg":  "ok",
		"files": gin.H{
			"file":    os.Getenv("IMGADDRESS")+ name,
			"imgName": name,
		},
	})

}
