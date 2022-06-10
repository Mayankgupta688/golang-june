package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/sessions"
)

func GetJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)
	keyValues := token.Claims.(jwt.MapClaims)
	keyValues["authorized"] = true
	keyValues["client"] = "Krissanawat"
	keyValues["aud"] = "billing.jwtgo.io"
	keyValues["iss"] = "jwtgo.io"
	keyValues["exp"] = time.Now().Add(time.Minute * 1).Unix()
	tokenString, _ := token.SignedString([]byte("ANY_SECRET_KEY"))
	return tokenString, nil
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":4000", nil)
}

var (
	key          = []byte("randon-encryption-key")
	sessionStore = sessions.NewCookieStore(key)
)

func login(w http.ResponseWriter, r *http.Request) {
	cookieData, _ := sessionStore.Get(r, "cookie-name")
	cookieData.Values["authenticated"] = true
	cookieData.Values["name"] = "Mayank"

	cookieData.Values["token"], _ = GetJWT()

	cookieData.Save(r, w)
	fmt.Fprintf(w, "Login")
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookieData, _ := sessionStore.Get(r, "cookie-name")
	cookieData.Values["authenticated"] = false
	cookieData.Values["name"] = "Unknown"
	cookieData.Values["token"] = ""
	cookieData.Save(r, w)
	fmt.Fprintf(w, "Logout")
}

func secret(w http.ResponseWriter, r *http.Request) {
	cookieData, _ := sessionStore.Get(r, "cookie-name")

	token, _ := jwt.Parse(fmt.Sprintf("%v", cookieData.Values["token"]), func(token *jwt.Token) (interface{}, error) {
		return "", nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	fmt.Println(claims["client"])

	if cookieData.Values["authenticated"] == false {
		fmt.Fprintf(w, "The User is not Authenticated: "+fmt.Sprintf("%v", cookieData.Values["name"]))
	} else {
		fmt.Fprintf(w, "The User is Authenticated: "+fmt.Sprintf("%v", cookieData.Values["name"]))
	}
}
