package controllers

import (
	"CamexAPI/models"
	u "CamexAPI/utils"
	"encoding/json"
	"net/http"
)

var CreateData = func(w http.ResponseWriter, r *http.Request) {
	imei := r.Context().Value("imei") . (string)
	data := &models.Data{ Imei: imei }

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		u.Respond(w, http.StatusForbidden, u.Message(false, "Error while decoding request body"))
		return
	}

	resp := data.Create()
	if resp["status"] . (bool) {
		u.Respond(w, http.StatusOK, resp)
	} else {
		u.Respond(w, http.StatusForbidden, resp)
	}
}

var GetDatas = func(w http.ResponseWriter, r *http.Request) {
	imei := r.Context().Value("imei") . (string)
	datas := models.GetDatas(imei)
	if datas == nil {
		u.Respond(w, http.StatusForbidden, u.Message(false, "Data is not recognized"))
		return
	}
	resp := u.Message(true, "Success")
	resp["data"] = datas
	u.Respond(w, http.StatusOK, resp)
}