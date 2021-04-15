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
	querystring := `
select ut.user_name,ut."password",ur.user_role, ut.email , ut.mob_no from user_table ut
inner join user_role ur on ur.role_id = ut.user_role`
	rows, err := pool.Query(context.Background(),querystring)
	if err != nil{
		return nil, err
	}
	for rows.Next(){
		var entity entity.UserExcel
		err :=rows.Scan(&entity.UserName,&entity.Password,&entity.UserRole,&entity.Email,&entity.MobNo)
		if err != nil{
			return nil, err
		}
		responce = append(responce, entity)
	}
	return responce, nil

}
