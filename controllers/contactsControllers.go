package controllers

import (
	"cashgone/models"
	u "cashgone/utils"
	"encoding/json"
	"net/http"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value("user") . (uint) //Grab the id of the user that send the request
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, http.StatusForbidden, u.Message(false, "Error while decoding request body"))
		return
	}

	contact.UserId = user
	resp := contact.Create()
	if resp["status"] == false {
		u.Respond(w, http.StatusForbidden, resp)
	} else if resp["status"] == true {
		u.Respond(w, http.StatusOK, resp)
	}

}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	user_id := r.Context().Value("user") . (uint)
	data := models.GetContacts(user_id)
	if data == nil {
		u.Respond(w, http.StatusForbidden, u.Message(false, "User is not recognized"))
		return
	}
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, http.StatusOK, resp)
}

//var GetContact = func(w http.ResponseWriter, r *http.Request) {
//	user_id := r.Context().Value("user") . (uint)
//
//}