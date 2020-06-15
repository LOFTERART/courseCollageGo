package admin

import (
	"PC/models"
	"PC/serializer"
	"PC/utils"
	"github.com/gin-gonic/gin"
)

type Admin struct {
	Name     string `form:"name" json:"name" `
	Password string `form:"passWord" json:"passWord" `
	Tokens   string `form:"tokens" json:"tokens" `
}

func (item *Admin) CreateAdmin() serializer.Response {

	admin := models.Admin{
		Name:     item.Name,
		Password: item.Password,
		Tokens:   utils.RandStringRunes(10),
	}

	//查询数据库是否存在

	B := models.DB.Where("name = ?", item.Name).First(&admin).RecordNotFound()

	if !B {
		return serializer.Response{
			Code:  20001,
			Data:  1,
			Msg:   "用户名已存在 重新输入用户名",
			Error: "",
		}
	} else {
		models.DB.Where(models.Admin{Name: item.Name}).FirstOrCreate(&admin)

		return serializer.Response{
			Code: 20000,
			Data: serializer.BuildAdminSerializer(admin),
		}

	}

}

//登录
func (item *Admin) LoginAdmin() serializer.Response {

	admin := models.Admin{
		Name:     item.Name,
		Password: item.Password,
	}

	//查询数据库是否存在

	B := models.DB.Where("name = ? AND password =? ", item.Name, item.Password).First(&admin).RecordNotFound()

	if B {
		return serializer.Response{
			Code: 20001,
			Msg:  "请检查用户名或者密码",
		}
	} else {
		return serializer.Response{
			Code: 20000,
			Data: gin.H{
				"token": admin.Tokens,
			},
		}

	}

}

//获取用户信息
func (item *Admin) GetAdmin() serializer.Response {

	var admin models.Admin
	//查询数据库是否存在

	models.DB.Where("tokens = ? ", item.Tokens).First(&admin)
	return serializer.Response{
		Code: 20000,
		Data: serializer.BuildAdminSerializer(admin),
	}

}
