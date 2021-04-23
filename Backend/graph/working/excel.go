package working

import (
	_ "image/gif"
	"log"
	"strconv"

	_ "image/jpeg"

	_ "image/png"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/shubhacker/gqlgen-todos/graph/entity"
)

func CreateExcel(entity []entity.UserExcel) string {
	f := excelize.NewFile()
	index := f.NewSheet("sheet 1")
	excelize.TitleToNumber("AK")
	f.SetCellValue("sheet 1", "A"+strconv.Itoa(1), "UserName")
	f.SetCellValue("sheet 1", "B"+strconv.Itoa(1), "Password")
	f.SetCellValue("sheet 1", "C"+strconv.Itoa(1), "UserRole")
	f.SetCellValue("sheet 1", "D"+strconv.Itoa(1), "Email")
	f.SetCellValue("sheet 1", "E"+strconv.Itoa(1), "MobileNo")
	op := 2
	for _, data := range entity {
		userColumn := "A" + strconv.Itoa(op)
		passwordColumn := "B" + strconv.Itoa(op)
		userRoleColumn := "C" + strconv.Itoa(op)
		emailColumn := "D" + strconv.Itoa(op)
		MobNoColumn := "E" + strconv.Itoa(op)
		f.SetCellValue("sheet 1", userColumn, data.UserName)
		f.SetCellValue("sheet 1", passwordColumn, data.Password)
		f.SetCellValue("sheet 1", userRoleColumn, data.UserRole)
		f.SetCellValue("sheet 1", emailColumn, data.Email)
		f.SetCellValue("sheet 1", MobNoColumn, data.MobNo)
		op += 1
	}
	f.SetActiveSheet(index)

	if err := f.SaveAs("Test.xlsx"); err != nil {
		log.Println(err.Error())
	}
	return "Excel Created"
}
func CreateExcelForMaster(entity []entity.MasterExcel) string {
	log.Println("CreateExcelForMaster()")
	f := excelize.NewFile()
	op := 10
	key1 := 0
	key2 := 0
	key3 := 0
	Column1 := "A" + strconv.Itoa(op)
	Column2 := "B" + strconv.Itoa(op)
	Column3 := "C" + strconv.Itoa(op)
	Column4 := "D" + strconv.Itoa(op)
	Column5 := "E" + strconv.Itoa(op)
	Column6 := "F" + strconv.Itoa(op)

	Username := make(map[int]string)
	Password := make(map[int]string)
	UserRole := make(map[int]string)
	Email := make(map[int]string)
	Mob := make(map[int]int)
	VideoTopic := make(map[int]string)
	VideoLink := make(map[int]string)
	Tool := make(map[int]string)
	ToolLink := make(map[int]string)
	Book := make(map[int]string)
	BookLink := make(map[int]string)
	BlogText := make(map[int]string)
	Reference := make(map[int]string)
	UserSheet := f.NewSheet("User Sheet")
	VideoSheet := f.NewSheet("Video Sheet")
	BlogSheet := f.NewSheet("Blog Sheet")
	err := f.AddPicture("User Sheet", "A1", "graph\\Images\\ExcelImage.png", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		log.Println("Error In User Picture", err)
	}
	err = f.AddPicture("Video Sheet", "A1", "graph\\Images\\ExcelImage.png", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		log.Println("Error In Video Picture", err)
	}
	err = f.AddPicture("Blog Sheet", "A1", "graph\\Images\\ExcelImage.png", `{"x_scale": 0.5, "y_scale": 0.5}`)
	if err != nil {
		log.Println("Error In Blog Picture", err)
	}
	f.SetCellValue("User Sheet", Column1, "UserName")
	f.SetCellValue("User Sheet", Column2, "PassWord")
	f.SetCellValue("User Sheet", Column3, "UserRole")
	f.SetCellValue("User Sheet", Column4, "Email")
	f.SetCellValue("User Sheet", Column5, "Mobile No")

	f.SetCellValue("Video Sheet", Column1, "Video Topic")
	f.SetCellValue("Video Sheet", Column2, "Video Link")
	f.SetCellValue("Video Sheet", Column3, "Book Name")
	f.SetCellValue("Video Sheet", Column4, "Book Link")
	f.SetCellValue("Video Sheet", Column5, "Tool Name")
	f.SetCellValue("Video Sheet", Column6, "Tool Link")

	f.SetCellValue("Blog Sheet", Column1, "Blog Text")
	f.SetCellValue("Blog Sheet", Column2, "Reference Link")

	for _, maping := range entity {
		if maping.UserName != "" {
			Username[key1] = maping.UserName
			Password[key1] = maping.Password
			UserRole[key1] = maping.UserRole
			Email[key1] = maping.Email
			Mob[key1] = maping.MobNo
			key1 += 1
		}
	}
	for _, maping := range entity {
		if maping.VideoTopic != "" {
			VideoTopic[key2] = maping.VideoTopic
			VideoLink[key2] = maping.VideoLink
			Tool[key2] = maping.ToolName
			ToolLink[key2] = maping.ToolLink
			Book[key2] = maping.BookName
			BookLink[key2] = maping.BookLink
			key2 += 1
		}
	}

	for _, maping := range entity {
		if maping.BlogText != "" {
			BlogText[key3] = maping.BlogText
			Reference[key3] = maping.ReferenceLink
			key3 += 1
		}

	}
	key1 = 0
	op = 13
	for _ = range Username {
		Column1 = "A" + strconv.Itoa(op)
		Column2 = "B" + strconv.Itoa(op)
		Column3 = "C" + strconv.Itoa(op)
		Column4 = "D" + strconv.Itoa(op)
		Column5 = "E" + strconv.Itoa(op)

		f.SetCellValue("User Sheet", Column1, Username[key1])
		f.SetCellValue("User Sheet", Column2, Password[key1])
		f.SetCellValue("User Sheet", Column3, UserRole[key1])
		f.SetCellValue("User Sheet", Column4, Email[key1])
		f.SetCellValue("User Sheet", Column5, Mob[key1])
		op += 1
		key1 += 1
	}
	key2 = 0
	op = 13
	for _ = range VideoTopic {
		Column1 = "A" + strconv.Itoa(op)
		Column2 = "B" + strconv.Itoa(op)
		Column3 = "C" + strconv.Itoa(op)
		Column4 = "D" + strconv.Itoa(op)
		Column5 = "E" + strconv.Itoa(op)
		Column6 = "F" + strconv.Itoa(op)

		f.SetCellValue("Video Sheet", Column1, VideoTopic[key2])
		f.SetCellValue("Video Sheet", Column2, VideoLink[key2])
		f.SetCellValue("Video Sheet", Column3, Book[key2])
		f.SetCellValue("Video Sheet", Column4, BookLink[key2])
		f.SetCellValue("Video Sheet", Column5, Tool[key2])
		f.SetCellValue("Video Sheet", Column6, ToolLink[key2])
		op += 1
		key2 += 1
	}
	key3 = 0
	op = 13
	for _ = range BlogText {
		Column1 = "A" + strconv.Itoa(op)
		Column2 = "B" + strconv.Itoa(op)
		f.SetCellValue("Blog Sheet", Column1, BlogText[key3])
		f.SetCellValue("Blog Sheet", Column2, Reference[key3])
		op += 1
		key3 += 1
	}

	f.SetActiveSheet(UserSheet)
	f.SetActiveSheet(VideoSheet)
	f.SetActiveSheet(BlogSheet)
	if err := f.SaveAs("Master.xlsx"); err != nil {
		log.Println(err.Error())
	}
	return "Excel Created Successfully"
}
