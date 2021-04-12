package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

var mySigningKey = []byte("mysupersecretphrase")

type contextkey struct{
	name string
}

func GenerateJWT(Username string, passWord string)(string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = Username
	claims["password"] = passWord
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString,err := token.SignedString(mySigningKey)
	if err != nil{
		log.Println("Something Went Wrong!",err)
	}
	return tokenString,err
}
