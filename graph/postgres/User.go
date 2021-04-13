package postgres

import (
	"context"
	"github.com/shubhacker/gqlgen-todos/graph/entity"
)

func FetchUserForExcel()([]entity.UserExcel, error){
	if pool == nil{
		pool = GetPool()
	}
	var responce []entity.UserExcel
	querystring := `select ut.user_name from user_table ut`
	rows, err := pool.Query(context.Background(),querystring)
	if err != nil{
		return nil, err
	}
	for rows.Next(){
		var entity entity.UserExcel
		err :=rows.Scan(&entity.UserName)
		if err != nil{
			return nil, err
		}
		responce = append(responce, entity)
	}
	return responce, nil

}
