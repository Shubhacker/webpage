package postgres

import (
	"context"
	"github.com/shubhacker/gqlgen-todos/graph/entity"
	"log"
)

func FetchMasterDataForExcel()[]entity.MasterExcel{
	log.Println("FetchMasterDataForExcel()")
	if pool == nil{
		pool = GetPool()
	}
	var responce []entity.MasterExcel
	querystringForUser := `select ut.user_name , ut."password" , ut.email ,ur.user_role , ut.email , ut.mob_no  
from user_table ut
inner join user_role ur on ur.role_id  = ut.user_role`
//	querystringForVideo := `select vt.video_topic, vt.video_link, b.book_name ,b.book_link ,t.tool_name ,t.tool_link  from video_table vt
//inner join book b on b.book_id = vt.book_id
//inner join tools t on t.tools_id  = vt.tools_id `
	rows, err := pool.Query(context.Background(), querystringForUser)
	if err != nil{
		log.Println("Error In Master Data Fetch")
	}
	//rows, err = pool.Query(context.Background(), querystringForVideo)
	//if err != nil{
	//	log.Println("Error In Master Data Fetch")
	//}
	for rows.Next(){
		var entity entity.MasterExcel
		err = rows.Scan(&entity.UserName,&entity.Password,&entity.Email, &entity.MobNo)
		responce = append(responce, entity)
	}
	log.Println("Inside--->",responce)
	return responce
}
