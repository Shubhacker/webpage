package postgres

import (
	"context"
	"log"

	"github.com/shubhacker/gqlgen-todos/graph/entity"
)

func FetchMasterDataForExcel() []entity.MasterExcel {
	log.Println("FetchMasterDataForExcel()")
	if pool == nil {
		pool = GetPool()
	}
	var responce []entity.MasterExcel
	querystringForUser := `select ut.user_name as username, ut."password" as passwords, ut.email as email,ur.user_role as userrole , ut.mob_no as mob  
	from user_table ut
	inner join user_role ur on ur.role_id  = ut.user_role`
	rowsForUser, err := pool.Query(context.Background(), querystringForUser)
	if err != nil {
		log.Println("Error In Master Data Fetch")
	}
	for rowsForUser.Next() {
		var entity entity.MasterExcel
		err = rowsForUser.Scan(&entity.UserName, &entity.Password, &entity.Email, &entity.UserRole, &entity.MobNo)
		if err != nil {
			log.Println(err.Error())
		}
		responce = append(responce, entity)
	}
	querystringForVideo := `select vt.video_topic as videotopic, vt.videolink as videolink, b.book_name as bookname,b.book_link as booklink,
	t.tool_name as tollname,t.tool_link as tollink  from video_table vt 
	inner join book b on b.book_id = vt.book_id 
	inner join tools t on t.tools_id  = vt.tools_id`
	rowsForVideo, err := pool.Query(context.Background(), querystringForVideo)
	if err != nil {
		log.Println("Error In Master Data Fetch")
	}
	for rowsForVideo.Next() {
		var entity entity.MasterExcel
		err = rowsForVideo.Scan(&entity.VideoTopic, &entity.VideoLink, &entity.BookName, &entity.BookLink, &entity.ToolName, &entity.ToolLink)
		if err != nil {
			log.Println(err.Error())
		}
		responce = append(responce, entity)
	}

	querystringForBlog := `select bg.blog_text, bg.reference_link from blog bg`
	rowsForBlog, err := pool.Query(context.Background(), querystringForBlog)
	if err != nil {
		log.Println("Error In Master Data Fetch")
	}
	for rowsForBlog.Next() {
		var entity entity.MasterExcel
		err = rowsForBlog.Scan(&entity.BlogText, &entity.ReferenceLink)
		if err != nil {
			log.Println(err.Error())
		}
		responce = append(responce, entity)
	}
	return responce
}
