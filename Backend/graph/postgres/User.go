package postgres

import (
	"context"
	"log"

	"github.com/shubhacker/gqlgen-todos/graph/entity"
)

func FetchUserForExcel() ([]entity.UserExcel, error) {
	if pool == nil {
		pool = GetPool()
	}
	var responce []entity.UserExcel
	querystring := `
select ut.user_name,ut."password",ur.user_role, ut.email , ut.mob_no from user_table ut
inner join user_role ur on ur.role_id = ut.user_role`
	rows, err := pool.Query(context.Background(), querystring)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var entity entity.UserExcel
		err := rows.Scan(&entity.UserName, &entity.Password, &entity.UserRole, &entity.Email, &entity.MobNo)
		if err != nil {
			return nil, err
		}
		responce = append(responce, entity)
	}
	return responce, nil

}

func FetchUserForLogin(UserName string) (entity.UserData, error) {
	if pool == nil {
		pool = GetPool()
	}
	var entity entity.UserData
	querystring := `select ut.user_name , ut.email , ut.mob_no ,ur.user_role from user_table ut 
	inner join user_role ur on ur.role_id = ut.user_role 
	where ut.user_name =$1`
	err := pool.QueryRow(context.Background(), querystring, UserName).Scan(&entity.UserName, &entity.Email, &entity.MobNo, &entity.UserRole)
	if err != nil {
		log.Println("ERROR", err.Error())
	}
	return entity, nil
}
