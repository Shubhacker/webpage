package entity

type UserExcel struct {
	UserName string
	Password string
	UserRole string
	Email    string
	MobNo    int
}

type MasterExcel struct {
	UserName      string
	Password      string
	UserRole      string
	Email         string
	MobNo         int
	BlogText      string
	ReferenceLink string
	BookName      string
	BookLink      string
	ToolName      string
	ToolLink      string
	VideoLink     string
	VideoTopic    string
}

type UserData struct {
	UserName string
	Email    string
	MobNo    int
	UserRole string
}
