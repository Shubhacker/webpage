package main

import (
	"log"
	"time"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mysupersecretphrase")

func GenerateJWT()(string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Piyush Takale"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString,err := token.SignedString(mySigningKey)
	if err != nil{
		log.Println("Something Went Wrong!",err)
	}
	return tokenString,err
}

func main(){
	log.Println("Creating JWT")
	tokenstring, err := GenerateJWT()
	if err != nil{
		log.Println("Error in creating JWT",err)
	}
	log.Println(tokenstring)
}