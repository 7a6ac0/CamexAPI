package controllers

import (
	"CamexAPI/models"
	u "CamexAPI/utils"
	"net/http"
)

var GetToken = func(w http.ResponseWriter, r *http.Request) {
	imei := r.FormValue("imei")
	token := &models.Token{ Imei: imei }

	resp := token.CreateToken()
	if resp["status"] . (bool) {
		u.Respond(w, http.StatusOK, resp)
	} else {
		u.Respond(w, http.StatusForbidden, resp)
	}
}

