package app

func init() {
	SetUpConf()
	SetUpBasics()
	SetUpDB()
	SetUpResult()
}

func Run() {
	Result.Running()
}

func Close() {
	//_ = Cache.Close()
	_ = DB.Close()
}
