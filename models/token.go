package models

import (
	u "CamexAPI/utils"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

/*
JWT claims struct
*/
type Token struct {
	Imei string `json:"imei"`
	jwt.StandardClaims
}

func (token *Token) CreateToken() (map[string]interface{}) {
	//Create JWT token
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"imei": token.Imei,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, _ := tk.SignedString([]byte(os.Getenv("token_password")))

	resp := u.Message(true, "Success")
	resp["token"] = tokenString
	return resp
}