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
//
////a struct to rep user account
//type Account struct {
//	gorm.Model
//	Email string `json:"email"`
//	Password string `json:"password"`
//	Token string `json:"token";sql:"-"`
//}

//Validate incoming user details...
//func (account *Account) Validate() (map[string] interface{}, bool) {
//
//	if !strings.Contains(account.Email, "@") {
//		return u.Message(false, "Email address is required"), false
//	}
//
//	if len(account.Password) < 6 {
//		return u.Message(false, "Password is required"), false
//	}
//
//	//Email must be unique
//	temp := &Account{}
//
//	//check for errors and duplicate emails
//	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return u.Message(false, "Connection error. Please retry"), false
//	}
//	if temp.Email != "" {
//		return u.Message(false, "Email address already in use by another user."), false
//	}
//
//	return u.Message(false, "Requirement passed"), true
//}

//func (account *Account) Create() (map[string] interface{}) {
//
//	if resp, ok := account.Validate(); !ok {
//		return resp
//	}
//
//	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
//	account.Password = string(hashedPassword)
//
//	GetDB().Create(account)
//
//	if account.ID <= 0 {
//		return u.Message(false, "Failed to create account, connection error.")
//	}
//
//	response := u.Message(true, "Account has been created")
//	return response
//}

func (token *Token) CreateToken() (map[string]interface{}) {
	//Create JWT token
	tk := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"imei": token.Imei,
		"exp": time.Now().Add(time.Hour * 2).Unix(),
	})
	tokenString, _ := tk.SignedString([]byte(os.Getenv("token_password")))

	resp := u.Message(true, "Get Token")
	resp["token"] = tokenString
	return resp
}

//func GetUser(u uint) *Account {
//
//	acc := &Account{}
//	GetDB().Table("accounts").Where("id = ?", u).First(acc)
//	if acc.Email == "" { //User not found!
//		return nil
//	}
//
//	acc.Password = ""
//	return acc
//}