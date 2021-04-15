package auth

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shubhacker/gqlgen-todos/graph/postgres"
	"log"
	"net/http"
	"strings"
	"time"
)

var mySigningKey = []byte("mysupersecretphrase")

type contextKey struct {
	name string
}

var userNameCtxKey = &contextKey{"user"}
var userRoleCtxKey = &contextKey{"UserRole"}

func GenerateJWT(Username string, Password string, UserRole string)(string, error){
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = Username
	claims["password"] = Password
	claims["UserRole"] = UserRole
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString,err := token.SignedString(mySigningKey)
	if err != nil{
		log.Println("Something Went Wrong!",err)
	}
	return tokenString,err
}

func AuthMiddleware(pool *pgxpool.Pool) func(handler http.Handler)http.Handler{
	return func(next http.Handler)http.Handler{
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request){
			header := request.Header.Get("Authorization")
			if header == ""{
				if request.RequestURI == "/v2/query"{
					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusUnauthorized)
				}
				next.ServeHTTP(response,request)
				return
			}else{
				Uname, err := ValidateAndGetUserName(header)
				if err != nil{
					if request.RequestURI == "/v2/query" {
						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusUnauthorized)
					}
					log.Println(err)
				}else{
					userModel, err := postgres.GetUserDetails(pool,Uname)
					if err == nil && userModel != nil{
						ctx := context.WithValue(request.Context(), userNameCtxKey, userModel.UserName)
						ctxUserRole := context.WithValue(ctx, userRoleCtxKey, userModel.UserRole)
						request = request.WithContext(ctxUserRole)
					}
				}
				next.ServeHTTP(response,request)
			}
		})
	}
}

func ValidateAndGetUserName(header string)(string, error){
authheader := strings.Split(header," ")
var tokenString string
var UserName string
if len(authheader)> 1{
	tokenString = strings.Split(header," ")[1]
	token,_ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok{
			return nil, fmt.Errorf("There is Error")
		}
		return jwtSecret, nil
	})
	if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid{
		UserName = claims["user"].(string)
	}else{
		return "", fmt.Errorf("Token Secret key is invalid")
	}
}else{
	return "",fmt.Errorf("Token secret key invalid")
}
return UserName, nil
}

func ForContext(ctx context.Context)(*string, *string){
//var UserId *int
var UserName *string
var UserRole *string
//userIdValue, userIdOk := ctx.Value(userCtxKey).(int)
userNameValue, userNameOk := ctx.Value(userNameCtxKey).(string)
userRoleValue, userRoleOk := ctx.Value(userRoleCtxKey).(string)
if  userNameOk && userRoleOk {
	//UserId = &userIdValue
	UserName = &userNameValue
	UserRole = &userRoleValue
}else{
	log.Println("Error in ForContext")
}
return UserName, UserRole
}

func GetAuthRole(ctx context.Context) *string {
	var authRole *string
	authRoleIDValue, ok := ctx.Value(userRoleCtxKey).(string)
	if ok {
		authRole = &authRoleIDValue
	}
	return authRole
}