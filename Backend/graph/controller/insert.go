package controller

import (
	"context"
	"log"

	"github.com/shubhacker/gqlgen-todos/graph/auth"

	"github.com/shubhacker/gqlgen-todos/graph/entity"
	"github.com/shubhacker/gqlgen-todos/graph/mapper"
	"github.com/shubhacker/gqlgen-todos/graph/model"
	"github.com/shubhacker/gqlgen-todos/graph/postgres"
)

func FetchTableData(ctx context.Context) *model.Fetch {
	data := postgres.FetchTableDataIn()
	mapData := mapper.MapFetchData(data)
	return mapData
}

func FetchToolData(ctx context.Context, input *model.FetchToolsInput) *model.ToolResponceData {
	var responce model.ToolResponceData
	Filter := mapper.MapFilterForTools(input)
	FilterMap := mapper.FilterTool(input.Filter)
	data := postgres.FetchToolDataFromDb(Filter.ID, FilterMap.Filter, FilterMap.FilterColumn)
	mapData := mapper.MapFetchDataForTools(data)
	responce.Data = mapData
	return &responce
}

func UpsertToolData(ctx context.Context, input model.UpsertTool) *model.UpsertToolResponce {
	var responce model.UpsertToolResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpsertForTools(input)
	err := postgres.UpsertToolData(mapData)
	if err != nil {
		responce.Message = "Error Upserting Tools"
	} else {
		responce.Message = "Upsert tool Succefully"
	}
	return &responce
}

func UpsertbookData(ctx context.Context, input model.UpsertBook) *model.UpsertBookResponce {
	var responce model.UpsertBookResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpsertForBooks(input)
	err := postgres.UpsertbookData(mapData)
	if err != nil {
		responce.Message = "Error Upserting Tools"
	} else {
		responce.Message = "Upsert book Succefully"
	}
	return &responce
}

func UpsertVideoData(ctx context.Context, input model.UpsertVideo) *model.UpsertVideoResponce {
	var responce model.UpsertVideoResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpsertForVideo(input)
	err := postgres.UpsertVideo(mapData)
	if err != nil {
		responce.Message = "Error Upserting Video"
	} else {
		responce.Message = "Upsert Video Succefully"
	}
	return &responce
}

func UpdateVideoData(ctx context.Context, input model.UpdateVideo) *model.UpsertVideoResponce {
	var responce model.UpsertVideoResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpdateForVideo(input)
	err := postgres.UpdateVideo(mapData)
	if err != nil {
		responce.Message = "Error Updating Video"
	} else {
		responce.Message = "Updating Video Succefully"
	}
	return &responce
}

func UpdatebookData(ctx context.Context, input model.UpdateBook) *model.UpsertBookResponce {
	var responce model.UpsertBookResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpdateForBooks(input)
	err := postgres.UpdateBookData(mapData)
	if err != nil {
		responce.Message = "Error Updating Books"
	} else {
		responce.Message = "Update book Succefully"
	}
	return &responce
}

func UpdateToolsData(ctx context.Context, input model.UpdateTools) *model.UpsertToolResponce {
	var responce model.UpsertToolResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpdateForTools(input)
	err := postgres.UpdateToolData(mapData)
	if err != nil {
		responce.Message = "Error Updating Tools"
	} else {
		responce.Message = "Update tool Succefully"
	}
	return &responce
}

func FetchBookData(ctx context.Context, input *model.FetchBookInput) *model.BookResponce {
	var responce model.BookResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developers" || *AuthRole != "tester" {
		responce.Error = true
		responce.Message = "You Do Not Have Permission For Action"
		responce.Data = nil
		return &responce
	}
	FilterData := mapper.FilterBook(input.Filter)
	data := postgres.FetchBookDataFromDb(input, FilterData.Filter, FilterData.FilterColumn)
	mapData := mapper.MapFetchBookData(data)
	responce.Data = mapData
	return &responce
}

func UpsertUserData(ctx context.Context, input model.UserUpsert) *model.UserResponce {
	var responce model.UserResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpsertForUser(input)
	err := postgres.UpsertUserData(mapData)
	if err != nil {
		responce.Message = err.Error()
		return &responce
	} else {
		responce.Message = "Upsert User Succefully"
	}
	return &responce
}

func UpdateUserData(ctx context.Context, input model.UpdateUser) *model.UserResponce {
	log.Println("UpdateUserData()")
	var responce model.UserResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	var err error
	var mapData *entity.UpdateUser
	mapData, err = mapper.MapUpdateForUser(input)
	if err != nil {
		responce.Message = err.Error()
		return &responce
	}
	err = postgres.UpdateUserData(mapData)
	if err != nil {
		responce.Message = err.Error()
		return &responce
	} else {
		responce.Message = "Update User Succefully"
	}
	return &responce
}

func FetchVideoData(ctx context.Context, input *model.FetchVideoInput) *model.FetchVideoResponce {
	var responce model.FetchVideoResponce
	data := postgres.FetchVideoData(input)
	mapData := mapper.MapVideoFetchData(data)
	responce.Data = mapData
	return &responce
}

func UpsertBlogData(ctx context.Context, input model.UpserBlogData) *model.BlogResponce {
	var responce model.BlogResponce
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developer" {
		responce.Message = "Developer Permission Require For Action"
	}
	mapData := mapper.MapUpsertForBlog(input)
	err := postgres.UpsertBlogData(mapData)
	if err != nil {
		responce.Message = err.Error()
		return &responce
	} else {
		responce.Message = "Upsert User Succefully"
	}
	return &responce
}

func FetchBlogData(ctx context.Context, input *model.FetchBlogInput) *model.ResponceFetchBlog {
	var responce model.ResponceFetchBlog
	data := postgres.FetchBlogDataFromDb(input)
	mapData := mapper.MapFetchBlogData(data)
	responce.Data = mapData

	return &responce
}

func LoginApi(ctx context.Context, input *model.Login) *model.LoginResponce {
	mapLogin := mapper.MappingForLogin(input)
	//LoginPostgres,err := postgres.AuthenticateUser(mapLogin)
	// postgres.TokenRegeneration()
	AuthRole := postgres.AuthRoleForUser(mapLogin.UserName)
	check, err := auth.GenerateJWT(mapLogin.UserName, mapLogin.Password, AuthRole)
	if err != nil {
		log.Println("error in creating JWT Token!")
	}
	if err != nil {
		log.Println("Error in Authenticate!")
	}
	mapTest := mapper.MappingLogin(check)
	return mapTest
}

func MasterAPI(ctx context.Context) *model.MasterFetch {
	var responce model.MasterFetch
	AuthRole := auth.GetAuthRole(ctx)
	if *AuthRole != "developers" {
		responce.Error = true
		responce.Message = "Only Admin Can Access MasterAPI"
		return &responce
	}
	data := postgres.FetchMasterBlogDataFromDb()
	Map := mapper.MapForMaster(data)
	responce.Data = Map
	return &responce
}
