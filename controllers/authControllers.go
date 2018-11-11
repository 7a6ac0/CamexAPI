package controllers

import (
	"CamexAPI/models"
	u "CamexAPI/utils"
	"net/http"
)

//var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
//
//	account := &models.Account{}
//	err := json.NewDecoder(r.Body).Decode(account) //decode the request body into struct and failed if any error occur
//	if err != nil {
//		u.Respond(w, http.StatusForbidden, u.Message(false, "Invalid request"))
//		return
//	}
//
//	resp := account.Create() //Create account
//	if resp["status"] . (bool) {
//		u.Respond(w, http.StatusOK, resp)
//	} else {
//		u.Respond(w, http.StatusForbidden, resp)
//	}
//}

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

