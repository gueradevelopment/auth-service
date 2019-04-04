package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

type Token struct {
	Email string
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
	Token string `json:"token"`
}

type Response struct {
	Msg   string
	Valid bool
}

func getToken(data []byte) Account {
	var userData UserData
	json.Unmarshal(data, &userData)

	tk := &Token{Email: userData.Email}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("TOKEN_PASSWORD")))

	account := Account{}
	account.Token = tokenString //Store the token in the response
	account.Email = userData.Email
	return account
}

func validateToken(w http.ResponseWriter, r *http.Request) {
	tokenC, err := r.Cookie("token")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	host := os.Getenv("CLIENT_HOST")

	if err != nil {
		http.Redirect(w, r, host+"/login", 403)
		return
	}
	token := tokenC.Value

	emailC, err := r.Cookie("email")
	if err != nil {
		http.Redirect(w, r, host+"/login", 403)
		return
	}
	email := emailC.Value

	if token == "" || email == "" {
		http.Redirect(w, r, host+"/login", 403)
		return
	}

	tokenParsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("TOKEN_PASSWORD")), nil
	})

	if claims, ok := tokenParsed.Claims.(jwt.MapClaims); ok && tokenParsed.Valid {
		fmt.Println(claims["Email"])
		response := Response{"Token is valid", true}
		js, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	} else {
		http.Redirect(w, r, host+"/login", 403)
	}

}
