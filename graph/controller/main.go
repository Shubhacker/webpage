package controller

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

func InitCodes(dbPool *pgxpool.Pool) {
	// 	codeService := postgres.CodeService{Pool: dbPool}
	// 	var codeError error
	// 	codes, codeError := codeService.GetCodes()
	// 	if codeError != nil {
	// 		log.Println(codeError)
	// 	}
}
