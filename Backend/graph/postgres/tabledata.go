package postgres

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	"github.com/shubhacker/gqlgen-todos/graph/entity"
	"github.com/shubhacker/gqlgen-todos/graph/model"
	"log"
)

func FetchTableDataIn() entity.Fetch {
	log.Println("FetchTableDataIn()")
	var response entity.Fetch
	if pool == nil {
		pool = GetPool()
	}
	querystring := `insert into tools (tool_name, tool_link,date_created,is_active) VALUES 
	('tool_name 4', 'tool_link 4', (select current_timestamp), true);`

	err := pool.QueryRow(context.Background(), querystring).Scan(response.Name, response.ProjectName)
	if err != nil {
		log.Printf("%s - Error: %s", err.Error())
	}
	return response
}

func FetchToolDataFromDb(input *int,Filter string, FilterColumn string) []*entity.FetchToolData {
	log.Println("FetchToolDataFromDb()")
	var response []*entity.FetchToolData
	if pool == nil {
		pool = GetPool()
	}
	querystring := `select tool_name , tool_link , is_active from tools where is_active = true`
	var inputargs []interface{}
	if &input != nil && Filter != "" && FilterColumn != ""{
		if &input != nil {
			querystring = querystring + ` tools_id = ?`
			inputargs = append(inputargs, input)
		}
		if &Filter != nil {
			if Filter == "asc"{
				order := FilterColumn
				switch order{
				case "tool_name":
					querystring += ` order by tool_name asc`
				case "tool_link":
					querystring += ` order by tool_link asc`
				}
			}else if Filter == "desc"{
				order := FilterColumn
				switch order{
				case "tool_name":
					querystring += ` order by tool_name desc`
				case "tool_link":
					querystring += ` order by tool_link desc`
				}
			}
		}
	}
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	rows, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s", err.Error())
	}
	for rows.Next() {
		var entity entity.FetchToolData
		err = rows.Scan(&entity.Tool_name, &entity.Tool_link, &entity.Is_active)
		if err != nil {
			log.Println("%s - Error: %s here", err.Error())
		}
		response = append(response, &entity)
	}
	return response
}

func UpsertToolData(entity entity.Toolupsert) error {
	log.Println("UpsertToolData()")
	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	if err != nil {
		log.Printf("%s - Error: %s", UpsertToolData, err.Error())
		return err
	}
	defer tx.Rollback(context.Background())
	querystring := `insert into tools (tool_name, tool_link,date_created,is_active) VALUES 
	($1, $2, (select current_timestamp), $3)`
	_, err = tx.Exec(context.Background(), querystring, entity.Toolname, entity.Toollink, entity.Active)
	if err != nil {
		log.Printf("%s - Error: %s here")
	}
	txErr := tx.Commit(context.Background())
	if txErr != nil {
		log.Printf("%s - Error: %s here 2", UpsertToolData, txErr.Error())
		return err
	}
	return nil
}

func UpsertbookData(entity entity.Bookupsert) error {
	log.Println("UpsertbookData()")
	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	if err != nil {
		log.Printf("%s - Error: %s", UpsertbookData, err.Error())
		return err
	}
	defer tx.Rollback(context.Background())
	querystring := `insert into book(book_name, book_link, date_created, is_active) VALUES 
	($1, $2, (select current_timestamp), $3),`
	_, err = tx.Exec(context.Background(), querystring, entity.Bookname, entity.Booklink, entity.Active)
	if err != nil {
		log.Printf("%s - Error: %s")
	}
	txErr := tx.Commit(context.Background())
	if txErr != nil {
		log.Printf("%s - Error: %s", UpsertbookData, txErr.Error())
		return err
	}
	return nil
}

func UpdateBookData(entity entity.BookUpdate) error {
	log.Println("UpdateBookData()")
	if pool == nil {
		pool = GetPool()
	}

	querystring := `update book set `
	var inputargs []interface{}
	if entity.Bookname != "" {
		querystring += ` book_name = ?`
		inputargs = append(inputargs, entity.Bookname)
		querystring += ` ,`
	}
	if entity.Booklink != "" {
		querystring += ` book_link = ?`
		inputargs = append(inputargs, entity.Booklink)
		querystring += ` ,`
	}
	if &entity.Active != nil {
		querystring += ` is_active = ?`
		inputargs = append(inputargs, entity.Active)
	}
	querystring += `  where book_id = ?`
	inputargs = append(inputargs, entity.ID)

	log.Println("inputargs", inputargs)
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	_, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s")
	}
	return nil
}

func UpdateToolData(entity entity.ToolUpdate) error {
	log.Println("UpdateBookData()")
	if pool == nil {
		pool = GetPool()
	}

	querystring := `update Tools set `
	var inputargs []interface{}
	if entity.Toolname != "" {
		querystring += ` tool_name = ?`
		inputargs = append(inputargs, entity.Toolname)
		querystring += ` ,`
	}
	if entity.Toollink != "" {
		querystring += ` tool_link = ?`
		inputargs = append(inputargs, entity.Toollink)
		querystring += ` ,`
	}
	if &entity.Active != nil {
		querystring += ` is_active = ?`
		inputargs = append(inputargs, entity.Active)
	}
	querystring += `  where tools_id = ?`
	inputargs = append(inputargs, entity.ID)

	log.Println("inputargs", inputargs)
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	_, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s")
	}
	return nil
}

func UpsertVideo(entity entity.VideoData) error {
	log.Println("UpsertVideo()")
	if pool == nil {
		pool = GetPool()
	}

	querystring := `insert into Video_table (videolink,paid,video_topic,is_active,`
	var inputargs []interface{}
	if entity.BookName != "" {
		querystring += `book_id`
		querystring += ` ,`
	}
	if entity.ToolName != "" {
		querystring += ` tools_id`
		querystring += ` ,`
	}
	querystring += ` date_created) VALUES 
	(?, ?, ?,?,  `
	inputargs = append(inputargs, entity.Video_link, entity.Paid, entity.Video_Topic, entity.Active)
	if entity.BookName != "" {
		querystring += `(select book_id from book where book_name = ?)`
		inputargs = append(inputargs, entity.BookName)
		querystring += ` ,`
	}
	if entity.ToolName != "" {
		querystring += ` (select tools_id from tools where tool_name =  ?)`
		inputargs = append(inputargs, entity.ToolName)
		querystring += ` ,`
	}
	querystring += ` (select current_timestamp))`

	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	_, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s")
	}
	return nil
}

func UpdateVideo(entity entity.UpdateVideoData) error {
	log.Println("UpsertVideo()")
	if pool == nil {
		pool = GetPool()
	}

	querystring := `update Video_table set `
	var inputargs []interface{}
	if entity.Video_Topic != "" {
		querystring += ` video_topic= ?`
		inputargs = append(inputargs, entity.Video_Topic)
		querystring += ` ,`
	}
	if entity.Video_link != "" {
		querystring += ` videolink= ?`
		inputargs = append(inputargs, entity.Video_link)
		querystring += ` ,`
	}
	if entity.BookName != "" {
		querystring += ` book_id=(select book_id from book where book_name = ?)`
		inputargs = append(inputargs, entity.BookName)
		querystring += ` ,`
	}
	if entity.ToolName != "" {
		querystring += ` tools_id= (select tools_id from tools where tool_name =  ?)`
		inputargs = append(inputargs, entity.ToolName)
		querystring += ` ,`
	}
	if entity.Paid != nil {
		querystring += ` paid= ?`
		inputargs = append(inputargs, entity.Paid)
		querystring += ` ,`
	}
	querystring += `is_active = true where video_id= ?`
	log.Println("querystring-->", querystring)
	inputargs = append(inputargs, entity.ID)
	log.Println("inputargs->", inputargs)
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	_, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s")
	}
	return nil
}

func FetchBookDataFromDb(input *model.FetchBookInput,Filter string, FilterColumn string) []entity.FetchBook {
	log.Println("FetchBookDataFromDb()")
	var responce []entity.FetchBook
	if pool == nil {
		pool = GetPool()
	}
	var inputargs []interface{}
	querystring := `select book_name, book_link from book where is_active = true`
	if input != nil && Filter != "" && Filter != ""{
		if input.ID != nil {
			querystring += ` book_id = ?`
			inputargs = append(inputargs, input.ID)
		}
		if &Filter != nil{
			order := FilterColumn
			if Filter == "asc"{
				switch order{
				case "book_name":
					querystring += ` order by book_name asc`
				case "book_link":
					querystring += ` order by book_link asc`
				}
			}else if Filter == "desc"{
				switch order{
				case "book_name":
					querystring += ` order by book_name desc`
				case "book_link":
					querystring += ` order by book_name desc`
				}
			}
		}
	}

	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	rows, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s here", err.Error())
	}
	for rows.Next() {
		var entity entity.FetchBook
		err = rows.Scan(&entity.Bookname, &entity.Booklink)
		if err != nil {
			log.Println("%s - Error: %s here", err.Error())
		}
		responce = append(responce, entity)
	}
	return responce
}

func UpsertUserData(entity entity.UpsertUser) error {
	log.Println("UpsertUserData()")
	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	var inputargs []interface{}
	if err != nil {
		log.Printf("%s - Error: %s", UpsertbookData, err.Error())
		return err
	}
	defer tx.Rollback(context.Background())
	querystring := `insert into user_table(user_name,password,is_active, user_role`
	if entity.Email != "" {
		querystring += `, email`
	}
	if &entity.Mob_no != nil {
		querystring += `, mob_no`
	}
	querystring += `,date_created)values
	(?, ?, ?,(select role_id from user_role where user_role = ? and is_active = true)`
	inputargs = append(inputargs, entity.User_name)
	inputargs = append(inputargs, entity.Password)
	inputargs = append(inputargs, entity.Is_active)
	inputargs = append(inputargs, entity.User_role)
	if entity.Email != "" {
		querystring += `,?`
		inputargs = append(inputargs, entity.Email)
	}
	if &entity.Mob_no != nil {
		querystring += `, ?`
		inputargs = append(inputargs, entity.Mob_no)
	}
	querystring += `, (select current_timestamp))`
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	_, err = tx.Exec(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s here")
		return err
	}
	txErr := tx.Commit(context.Background())
	if txErr != nil {
		log.Printf("%s - Error: %s here 2", UpsertbookData, txErr.Error())
		return err
	}
	return nil
}

func IsPasswordRight(UserName string, Password string) bool {
	functionName := "IsPasswordRight()"
	log.Println(functionName)
	if pool == nil {
		pool = GetPool()
	}
	var result bool
	querystring := `select 1 from user_table ut where ut.user_name = $1 and ut."password" = $2;`
	var hasValue int
	err := pool.QueryRow(context.Background(), querystring, UserName, Password).Scan(&hasValue)
	if err == nil {
		result = true
	}
	return result
}

func UpdateUserData(entity *entity.UpdateUser) error {
	log.Println("UpdateUserData()")
	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	var inputargs []interface{}
	if err != nil {
		log.Printf("%s - Error: %s", UpsertbookData, err.Error())
		return err
	}
	defer tx.Rollback(context.Background())
	querystring := `update user_table set`
	if entity.Email != "" {
		querystring += ` email = ?`
		inputargs = append(inputargs, entity.Email)
		querystring += `, `
	}
	if entity.Mob_no != nil {
		querystring += ` mob_no = ?`
		inputargs = append(inputargs, entity.Mob_no)
		querystring += `, `

	}
	if entity.Password != "" {
		querystring += ` "password"= ?`
		inputargs = append(inputargs, entity.Password)
		querystring += `, `

	}
	if entity.User_role != "" {
		querystring += ` user_role = (select role_id from user_role where user_role = ? and is_active = true)`
		inputargs = append(inputargs, entity.User_role)
		querystring += `, `

	}
	if entity.Is_active != nil {
		querystring += `is_active = ?`
		inputargs = append(inputargs, entity.Is_active)
		querystring += `, `
	}
	querystring += ` date_modified = (select current_timestamp) where user_name = ?;`
	inputargs = append(inputargs, entity.User_name)
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	_, err = tx.Exec(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s here")
		return err
	}
	txErr := tx.Commit(context.Background())
	if txErr != nil {
		log.Printf("%s - Error: %s here 2", UpsertbookData, txErr.Error())
		return err
	}
	return nil
}

func FetchVideoData(input *model.FetchVideoInput) []*entity.FetchVideoData {
	log.Println("FetchVideoData()")
	var response []*entity.FetchVideoData
	if pool == nil {
		pool = GetPool()
	}
	querystring := `select vt.videolink , vt.paid , vt.is_active , vt.video_topic  , b.book_name , t.tool_name from video_table vt
	inner join book b on b.book_id = vt.book_id 
	inner join tools t on t.tools_id = vt.tools_id 
	where b.is_active = true and t.is_active = true `
	var inputargs []interface{}
	if input != nil{
		if input.VideoID != nil {
			querystring = querystring + ` and vt.video_id = ?`
			inputargs = append(inputargs, input.VideoID)
		}
		if input.VideoTopic != nil {
			querystring = querystring + ` and vt.video_topic = ?`
			inputargs = append(inputargs, input.VideoTopic)
		}
		if input.Paid != nil {
			querystring = querystring + ` and vt.paid = ?`
			inputargs = append(inputargs, input.Paid)
		}
	}
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	rows, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s", err.Error())
	}
	for rows.Next() {
		var entity entity.FetchVideoData
		err = rows.Scan(&entity.Video_link, &entity.Paid, &entity.Status, &entity.Video_Topic, &entity.BookName, &entity.ToolName)
		if err != nil {
			log.Println("%s - Error: %s here 2", err.Error())
		}
		response = append(response, &entity)
	}
	return response
}


func UpsertBlogData(entity entity.UpsertBlog) error {
	log.Println("UpsertUserData()")
	if pool == nil {
		pool = GetPool()
	}
	tx, err := pool.Begin(context.Background())
	var inputargs []interface{}
	if err != nil {
		log.Printf("%s - Error: %s", UpsertbookData, err.Error())
		return err
	}
	defer tx.Rollback(context.Background())
	querystring := `insert into Blog(blog_text`
	if entity.Videotopic != "" {
		querystring += `, video_id`
	}
	if entity.Bookname != "" {
		querystring += `, book_id`
	}
	if entity.Toolname != "" {
		querystring += `, tools_id`
	}
	if entity.Referencelink != "" {
		querystring += `, reference_link`
	}
	if &entity.Status != nil {
		querystring += `, is_active`
	}
	querystring += `, date_created) Values (?`
	inputargs = append(inputargs, entity.BlogText)

	if entity.Videotopic != "" {
		querystring += `,(select video_id from Video_table where video_topic = ?)`
		inputargs = append(inputargs, entity.Videotopic)
	}
	if entity.Bookname != "" {
		querystring += `, (select book_id from book where book_name = ?)`
		inputargs = append(inputargs, entity.Bookname)
	}
	if entity.Toolname != "" {
		querystring += `, (select tools_id from tools where tool_name = ?)`
		inputargs = append(inputargs, entity.Toolname)
	}
	if entity.Referencelink != "" {
		querystring += `, ?`
		inputargs = append(inputargs, entity.Referencelink)
	}
	if &entity.Status != nil {
		querystring += `, ?`
		inputargs = append(inputargs, entity.Status)
	}
	querystring += `,(select current_timestamp))`
	querystring = sqlx.Rebind(sqlx.DOLLAR, querystring)
	_, err = tx.Exec(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s here")
		return err
	}
	txErr := tx.Commit(context.Background())
	if txErr != nil {
		log.Printf("%s - Error: %s here 2", UpsertbookData, txErr.Error())
		return err
	}
	return nil
}


func FetchBlogDataFromDb(input *model.FetchBlogInput) []entity.FetchBlogData {
	log.Println("FetchBookDataFromDb()")
	var inputargs []interface{}
	var responce []entity.FetchBlogData
	if pool == nil {
		pool = GetPool()
	}
	querystring := `select b.blog_text , vt.video_topic , b2.book_name , t.tool_name , b.is_active , b.reference_link from blog b
inner join book b2 on b2.book_id = b.book_id 
inner join tools t on t.tools_id = b.tools_id 
inner join video_table vt on vt.video_id = b.video_id`
if input != nil{
	querystring += ` where b.blog_id = $1`
	inputargs = append(inputargs, input.BlogID)
}
	rows, err := pool.Query(context.Background(), querystring, inputargs...)
	if err != nil {
		log.Printf("%s - Error: %s here", err.Error())
	}
	for rows.Next() {
		var entity entity.FetchBlogData
		err = rows.Scan(&entity.BlogText, &entity.Videotopic, &entity.Bookname, &entity.Toolname, &entity.Status, &entity.Referencelink)
		if err != nil {
			log.Println("%s - Error: %s here 2", err.Error())
		}
		responce = append(responce, entity)
	}
	return responce
}


func FetchMasterBlogDataFromDb() []entity.FetchBlog {
	log.Println("FetchBookDataFromDb()")
	var responce []entity.FetchBlog
	if pool == nil {
		pool = GetPool()
	}
	querystring := `select b.blog_id ,b.blog_text , vt.video_id ,vt.video_topic , b2.book_id ,b2.book_name , t.tools_id, t.tool_name from blog b
inner join book b2 on b2.book_id = b.book_id
inner join tools t on t.tools_id = b.tools_id
inner join video_table vt on vt.video_id = b.video_id`
	rows, err := pool.Query(context.Background(), querystring)
	if err != nil {
		log.Printf("%s - Error: %s here", err.Error())
	}
	for rows.Next() {
		var entity entity.FetchBlog
		err = rows.Scan(&entity.BlogId, &entity.BlogText, &entity.VideoId, &entity.VideoTopic, &entity.BookId, &entity.BookName, &entity.ToolId,  &entity.ToolName)
		if err != nil {
			log.Println("%s - Error: %s here 2", err.Error())
		}
		responce = append(responce, entity)
	}
	return responce
}


func UserCheck() *entity.UserCheck {
	log.Println("FetchBookDataFromDb()")
	var entity entity.UserCheck
	if pool == nil {
		pool = GetPool()
	}
	querystring := `select ut.user_name, ut."password" , ur.user_role from user_table ut 
inner join user_role ur on ur.role_id = ut.user_role 
where ut.user_name = 'shubham'`
	 err := pool.QueryRow(context.Background(), querystring).Scan(&entity.UserName,&entity.Password ,&entity.UserRole)
	if err != nil {
		log.Printf("%s - Error: %s here", err.Error())
	}
	return &entity
}

//func AuthenticateUser(userEntity *entity.Login) (*entity.LoginResponce, error) {
//	var entity entity.LoginResponce
//	AuthRole := AuthRoleForUser(userEntity.UserName)
//	check,err := auth.GenerateJWT(userEntity.UserName, userEntity.Password,AuthRole)
//	if err!= nil{
//		log.Println("error in creating JWT Token!")
//	}
//	entity.JwtToken = &check
//	return &entity, nil
//}

func GetUserDetails(pool *pgxpool.Pool, UserName string) (*entity.UserDetails, error) {
	var userModel entity.UserDetails
	if pool != nil {
		queryString := `select ut.user_id ,ut.user_name ,ur.user_role from user_table ut 
inner join user_role ur on ur.role_id = ut.user_role 
where ut.user_name = $1`
		err := pool.QueryRow(context.Background(), queryString, UserName).Scan(&userModel.UserID, &userModel.UserName, &userModel.UserRole)
		if err != nil {
			return nil, err
		}
	}
	return &userModel, nil
}