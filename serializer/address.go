package serializer

import "PC/models"

type Address struct {
	Id uint `json:"id"`
	//省市县
	Province string `json:"province"`
	City     string `json:"city"`
	County   string `json:"county"`
	//详细地址
	FullAddress string `json:"full_address"`
	//	真实姓名
	TrueName string `json:"true_name"`
	TelPhone string `json:"tel_phone"`
	//拼接省市县
	FullAddressAdd string`json:"full_address_add"`
}

func FormatAddress(addr models.Address) *Address {
	return &Address{
		Id:          addr.ID,
		Province:    addr.Province,
		City:        addr.City,
		County:      addr.County,
		FullAddress: addr.FullAddress,
		FullAddressAdd: addr.Province+addr.City+addr.County+addr.County+addr.FullAddress,
		TrueName:    addr.TrueName,
		TelPhone:    addr.TelPhone,
	}
}

func FormatAddressList(addr []*models.Address) (addrlist []*Address) {
	for _, v := range addr {
		addrlist = append(addrlist, FormatAddress(*v))
	}

	return

}
