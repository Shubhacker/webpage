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
)

var jwtSecret = []byte("mysupersecretphrase")

type contextKey struct {
	name string
}

var userCtxKey = &contextKey{"UserID"}

func validateAndGetUserADName(header string) (string, error) {
	authHeader := strings.Split(header, " ")
	var tokenString string
	var adName string
	if len(authHeader) > 1 {
		tokenString = strings.Split(header, " ")[1]
		// Parse Token
		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return jwtSecret, nil
		})

		//Verify token secret
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			adName = claims["loggedInAs"].(string)

		} else {
			return "", fmt.Errorf("Token secret key is invalid!")
		}

	} else {
		return "", fmt.Errorf("Token secret key is invalid!")
	}

	return adName, nil
}


func AuthMiddleWare(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			header := request.Header.Get("Authorization")
			if header == "" {
				if request.RequestURI == "/v2/query" {
					response.Header().Set("Content-Type", "application/json")
					response.WriteHeader(http.StatusUnauthorized)
				}
				next.ServeHTTP(response, request)
				return
			} else {
				adName, err := validateAndGetUserADName(header)
				if err != nil {
					if request.RequestURI == "/v2/query" {
						response.Header().Set("Content-Type", "application/json")
						response.WriteHeader(http.StatusUnauthorized)
					}
					log.Println(err)
				} else {
					userModel := postgres.UserCheck(adName)
					if userModel != nil {
						log.Println("adname--->", adName)
						ctx := context.WithValue(request.Context(), userCtxKey, userModel.UserId)
						ctxSalesorg := context.WithValue(ctx, userCtxKey, nil)
						request = request.WithContext(ctxSalesorg)
					}
				}
				next.ServeHTTP(response, request)
			}
		})

	}
}

func ForContext(ctx context.Context)(*string){
	var userID *string
	userIDValue, userOk := ctx.Value(userCtxKey).(string)
	if userOk {
		userID = &userIDValue
	}
	return userID
}