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
	querystringForUser := `with usedata as(
select ut.user_name as username, ut."password" as passwords, ut.email as email,ur.user_role as userrole , ut.mob_no as mob  
from user_table ut
inner join user_role ur on ur.role_id  = ut.user_role
), videodata as(
select vt.video_topic as videotopic, vt.video_link as videolink, b.book_name as bookname,b.book_link as booklink,
t.tool_name as tollname,t.tool_link as tollink  from video_table vt 
inner join book b on b.book_id = vt.book_id 
inner join tools t on t.tools_id  = vt.tools_id 
), blogdata as (
select bg.blog_text, bg.reference_link from blog bg
)
select * from usedata, videodata,blogdata`
	rows, err := pool.Query(context.Background(), querystringForUser)
	if err != nil{
		log.Println("Error In Master Data Fetch")
	}
	for rows.Next(){
		var entity entity.MasterExcel
		err = rows.Scan(&entity.UserName,&entity.Password,&entity.Email, &entity.UserRole,&entity.MobNo,&entity.VideoTopic,&entity.VideoLink,
		&entity.BookName,&entity.BookLink, &entity.ToolName,&entity.ToolLink,&entity.BlogText,&entity.ReferenceLink)
		if err!= nil{
			log.Println(err.Error())
		}
		responce = append(responce, entity)
	}
	log.Println("Inside--->",responce)
	return responce
}
