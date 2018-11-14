package models

import (
	u "CamexAPI/utils"
	"github.com/jinzhu/gorm"
)

type Data struct {
	gorm.Model
	DataId string `json:"data_id"`
	Path string `json:"path"`
	Md5 string `json:"md5"`
	Timestamp int64 `json:"timestamp"`
	Imei string `json:"imei"` //The user that this contact belongs to
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (data *Data) Validate() (map[string] interface{}, bool) {

	if data.DataId == "" {
		return u.Message(false, "Data ID should be on the payload"), false
	}

	if data.Path == "" {
		return u.Message(false, "Path should be on the payload"), false
	}

	if data.Md5 == "" {
		return u.Message(false, "Md5 should be on the payload"), false
	}

	if data.Imei == "" {
		return u.Message(false, "IMEI should be on the payload"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (data *Data) Create() (map[string] interface{}) {
	if resp, ok := data.Validate(); !ok {
		resp["data"] = data
		return resp
	}

	GetDB().Create(data)

	resp := u.Message(true, "success")
	resp["data"] = data
	return resp
}

//func GetData(userId uint, contactId string) (*Data) {
//	contact := &Data{}
//	if err := GetDB().Table("data").Where("user_id = ? and id = ?", userId, contactId).First(contact).Error;
//	err != nil {
//		return nil
//	}
//	return contact
//}

func GetDatas(imei string) ([]*Data) {
	datas := make([]*Data, 0)
	if err := GetDB().Table("data").Where("imei = ?", imei).Find(&datas).Error;
	err != nil {
		return nil
	}

	return datas
}

