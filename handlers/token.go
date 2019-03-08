package handlers

import (
	"encoding/json"
	"os"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
}

type UserData struct {
	ID         string `json:"id"`
	Email      string `json:"email"`
	Verified   string `json:"verified_email"`
	Name       string `json:"name"`
	GivenName  string `json:"given_name"`
	FamilyName string `json:"family_name"`
	PictureURL string `json:"picture"`
}

type Account struct {
	Email string `json:"email"`
	Token string `json:"token";sql:"-"`
}

func getToken(data []byte) Account {
	var userData UserData
	json.Unmarshal(data, &userData)
	idInt, _ := strconv.ParseInt(userData.ID, 10, 62)
	tk := &Token{UserId: uint(idInt)}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	account := Account{}
	account.Token = tokenString //Store the token in the response
	account.Email = userData.Email
	return account
}
