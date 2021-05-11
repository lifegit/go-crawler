package newServices

type Bro struct {
	ServiceName string
	ServiceDir  string `json:"app_dir" form:"app_dir"`
}
