package mapper

import (
	"errors"
	"log"

	"github.com/shubhacker/gqlgen-todos/graph/entity"
	"github.com/shubhacker/gqlgen-todos/graph/model"
	"github.com/shubhacker/gqlgen-todos/graph/postgres"
)

func MapFetchData(entity entity.Fetch) *model.Fetch {
	log.Println("MapFetchData()")
	var out model.Fetch
	out.Employeename = entity.Name
	out.Projectename = entity.ProjectName
	return &out
}

func MapFetchDataForTools(tools []*entity.FetchToolData) []*model.FetchToolData {
	var out []*model.FetchToolData

	for _, items := range tools {
		var bookentity model.FetchToolData
		bookentity.Toolname = items.Tool_name
		bookentity.Toollink = items.Tool_link
		bookentity.IsActive = items.Is_active
		out = append(out, &bookentity)
	}
	return out
}

func MapUpsertForTools(input model.UpsertTool) entity.Toolupsert {
	log.Println("MapUpsertForTools()")
	var data entity.Toolupsert
	data.Toolname = input.ToolName
	data.Toollink = input.ToolLink
	data.Active = input.Status
	return data
}

func MapUpsertForBooks(input model.UpsertBook) entity.Bookupsert {
	log.Println("MapUpsertForBooks()")
	var data entity.Bookupsert
	data.Bookname = input.BookName
	data.Booklink = input.BookLink
	data.Active = input.Status
	return data
}

func MapUpsertForVideo(input model.UpsertVideo) entity.VideoData {
	log.Println("MapUpsertForVideo()")
	var data entity.VideoData
	data.Video_link = input.VideoLink
	data.Video_Topic = input.VideoTopic
	data.Active = input.Status
	data.Paid = input.Paid
	if input.BookName != nil {
		data.BookName = *input.BookName
	}
	if input.ToolName != nil {
		data.ToolName = *input.ToolName
	}
	return data
}

func MapUpsertForUser(input model.UserUpsert) entity.UpsertUser {
	log.Println("MapUpsertForUser()")
	var data entity.UpsertUser
	data.User_name = input.UserName
	data.Password = input.Password
	data.Is_active = input.IsActive
	data.User_role = input.UserRole
	if input.Email != nil {
		data.Email = *input.Email
	}
	if input.MobNo != nil {
		data.Mob_no = *input.MobNo
	}
	return data
}

func MapUpdateForVideo(input model.UpdateVideo) entity.UpdateVideoData {
	log.Println("MapUpdateForVideo()")
	var data entity.UpdateVideoData
	data.ID = input.VideoID
	if input.VideoTopic != nil {
		data.Video_Topic = *input.VideoTopic
	}
	if input.VideoLink != nil {
		data.Video_Topic = *input.VideoLink
	}
	if input.BookName != nil {
		data.BookName = *input.BookName
	}
	if input.ToolName != nil {
		data.ToolName = *input.ToolName
	}
	if input.Paid != nil {
		data.Paid = input.Paid
	}
	if input.Status != nil {
		data.Active = input.Status
	}
	return data
}
func MapUpdateForBooks(input model.UpdateBook) entity.BookUpdate {
	log.Println("MapUpsertForBooks()")
	var data entity.BookUpdate
	data.ID = input.BookID
	if input.BookName != nil {
		data.Bookname = *input.BookName
	}
	if input.BookLink != nil {
		data.Booklink = *input.BookLink
	}
	if input.Status != nil {
		data.Active = *input.Status
	}

	return data
}

func MapUpdateForTools(input model.UpdateTools) entity.ToolUpdate {
	log.Println("MapUpdateForTools()")
	var data entity.ToolUpdate
	data.ID = input.ToolID
	if input.ToolName != nil {
		data.Toolname = *input.ToolName
	}
	if input.ToolLink != nil {
		data.Toollink = *input.ToolLink
	}
	if input.Status != nil {
		data.Active = *input.Status
	}

	return data
}

func MapFetchBookData(entity []entity.FetchBook) []*model.FetchBookResponce {
	log.Println("MapFetchBookData()")
	var out []*model.FetchBookResponce

	for _, items := range entity {
		var bookentity model.FetchBookResponce
		bookentity.BookName = items.Bookname
		bookentity.BookLink = items.Booklink
		out = append(out, &bookentity)
	}
	return out
}

func MapUpdateForUser(input model.UpdateUser) (*entity.UpdateUser, error) {
	log.Println("MapUpdateForUser()")
	var data entity.UpdateUser
	data.User_name = input.UserName
	if input.Password != nil {
		if input.OldPassword != nil {
			if postgres.IsPasswordRight(input.UserName, *input.OldPassword) {
				data.Password = *input.Password
			} else {
				return nil, errors.New("Old Password Is Wrong")
			}
		} else {
			return nil, errors.New("Provide Old Password")
		}
	}
	if &input.IsActive != nil {
		data.Is_active = input.IsActive
	}
	if input.UserRole != nil {
		data.User_role = *input.UserRole
	}
	if input.Email != nil {
		data.Email = *input.Email
	}
	if input.MobNo != nil {
		data.Mob_no = input.MobNo
	}
	return &data, nil
}

func MapVideoFetchData(entity []*entity.FetchVideoData) []*model.FetchVideo {
	log.Println("MapVideoFetchData()")
	var out []*model.FetchVideo

	for _, items := range entity {
		var bookentity model.FetchVideo
		bookentity.VideoTopic = items.Video_Topic
		bookentity.VideoLink = items.Video_link
		bookentity.Paid = *items.Paid
		bookentity.BookName = &items.BookName
		bookentity.ToolName = &items.ToolName

		out = append(out, &bookentity)
	}
	return out
}

func MapUpsertForBlog(input model.UpserBlogData) entity.UpsertBlog {
	log.Println("MapUpsertForBlog()")
	var data entity.UpsertBlog
	data.BlogText = input.BlogText
	if input.Bookname != nil{
		data.Bookname = *input.Bookname
	}
	if input.Referencelink != nil{
		data.Referencelink = *input.Referencelink
	}
	if input.Status != nil{
		data.Status = *input.Status
	}
	if input.Toolname != nil{
		data.Toolname = *input.Toolname
	}
	if input.Videotopic != nil{
		data.Videotopic = *input.Videotopic
	}
	return data
}

func MapFetchBlogData(input []entity.FetchBlogData) []*model.FetchBlog {
	log.Println("MapFetchBookData()")
	var out []*model.FetchBlog

	for _, items := range input {
		var bookentity model.FetchBlog
		bookentity.Videotopic = &items.Videotopic
		bookentity.Bookname = &items.Bookname
		bookentity.Toolname = &items.Toolname
		bookentity.Referencelink = &items.Referencelink
		bookentity.Status = &items.Status
		bookentity.BlogText = *items.BlogText
		out = append(out, &bookentity)
	}
	return out
}