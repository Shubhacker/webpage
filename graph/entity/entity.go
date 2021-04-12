package entity

type Fetch struct {
	Name        string
	ProjectName string
}

type Tool struct {
	Tool_name *string
	Tool_link *string
	Is_Active *bool
}

type Toolupsert struct {
	Toolname string
	Toollink string
	Active   bool
}

type ToolUpdate struct {
	ID       int
	Toolname string
	Toollink string
	Active   bool
}

type Bookupsert struct {
	Bookname string
	Booklink string
	Active   bool
}

type BookUpdate struct {
	ID       int
	Bookname string
	Booklink string
	Active   bool
}

type VideoData struct {
	Video_link  string
	Paid        bool
	Video_Topic string
	BookName    string
	ToolName    string
	Active      bool
}

type UpdateVideoData struct {
	ID          int
	Video_link  string
	Paid        *bool
	Video_Topic string
	BookName    string
	ToolName    string
	Active      *bool
}

type FetchBook struct {
	Bookname string
	Booklink string
}

type ToolFilter struct {
	Filter string
	FilterColumn string
}

type FetchBlogData struct{
	BlogText *string
	Videotopic string
	Bookname string
	Toolname string
	Referencelink string
	Status bool
}

type UserCheck struct{
	UserName string
	Password string
	UserRole string
}

type UserDetails struct{
	UserID int
	UserName string
	UserRole string
}

type Login struct {
	UserName string
	Password string
}

type LoginResponce struct{
	JwtToken *string
}

type FetchToolData struct {
	Tool_name *string
	Tool_link *string
	Is_active *bool
}

type FilterForTools struct {
	ID *int
}

type Credentials struct{
	Username string
	Password string
	AuthRole string
}

type FilterForBook struct{
	Filter string
	FilterColumn string
}

type UpsertUser struct {
	User_name string
	Password  string
	Email     string
	Mob_no    int
	User_role string
	Is_active bool
}

type UpdateUser struct {
	User_name   string
	OldPassword string
	Password    string
	Email       string
	Mob_no      *int
	User_role   string
	Is_active   *bool
}

type FetchVideoData struct {
	Video_link  string
	Paid        *bool
	Video_Topic string
	BookName    string
	ToolName    string
	Status      *bool
}

type UpsertBlog struct {
	BlogText      string
	Videotopic    string
	Bookname      string
	Toolname      string
	Referencelink string
	Status        bool
}
