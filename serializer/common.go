package serializer

import (
	"time"
)

// Response 基础序列化器
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Msg   string      `json:"msg"`
	Error string      `json:"error"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// BuildListResponse 列表构建器
func BuildListResponse(items interface{}, total uint, code int) Response {
	return Response{
		Data: DataList{
			Items: items,
			Total: total,
		},
		Code: code,
	}
}

func FormatTimeRFC(name time.Time) string {
	return name.Format("2006-01-02 15:04:05")
	//str:=name.Format("2006-01-02")
	//if str == "0001-01-01"{
	//	return ""
	//}
	//return str
}

//时间字符串=>时间戳
func FormatTimeStamp(name string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", name, loc)
	return tt.Unix()

}
