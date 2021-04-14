package controller

import (
	"github.com/shubhacker/gqlgen-todos/graph/auth"
	"github.com/shubhacker/gqlgen-todos/graph/model"
	"context"
	"github.com/shubhacker/gqlgen-todos/graph/postgres"
	"github.com/shubhacker/gqlgen-todos/graph/working"
	"log"
)

func CreateExcelForUser(ctx context.Context)*model.ExcelUserResponce{
	var responce model.ExcelUserResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole!= "developers"{
		responce.Error = true
		responce.Message = "Need Developer Permission"
		return &responce
	}
	GetData, err := postgres.FetchUserForExcel()
	if err != nil{
		responce.Error = true
		responce.Message = err.Error()
	}
	Excel := working.CreateExcel(GetData)
	responce.Error = false
	responce.Message = Excel
	return &responce
}


func MaterExcelFetch(ctx context.Context)*model.MasterExcelResponce{
	var responce model.MasterExcelResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developers"{
		responce.Error = true
		responce.Message = "Need Developer Permissiono!"
	}
	FetchMasterData := postgres.FetchMasterDataForExcel()
	//CreateExcel := working.CreateExcelForMaster()
	log.Println(FetchMasterData)
	responce.Error = false
	responce.Message = "Done"
	return &responce

}
