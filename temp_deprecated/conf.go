package temp

// use go-ini

// Deprecated

//type App struct {
//	JwtSecret       string
//	PageSize        int
//	RuntimeRootPath string
//
//	ImagePrefixUrl string
//	ImageSavePath  string
//	ImageMaxSize   int
//	ImageAllowExt  []string
//
//	LogSavePath string
//	LogSaveName string
//	LogFileExt  string
//	TimeFormat  string
//}
//
//var AppSetting = &App{}
//
//type Server struct {
//	RunMode      string
//	HttpPort     int
//	ReadTimeout  time.Duration
//	WriteTimeout time.Duration
//}
//
//var ServerSetting = &Server{}
//
//type Database struct {
//	User        string
//	Password    string
//	Host        string
//	Name        string
//	Port        int
//	TablePrefix string
//}
//
//var DatabaseSetting = &Database{}
//
//type Redis struct {
//	Host        string
//	Password    string
//	MaxIdle     int
//	MaxActive   int
//	IdleTimeout time.Duration
//}
//
//var RedisSetting = &Redis{}
//
//var Logger = zap.NewExample().Sugar()
//
//func SetUp() {
//	Cfg, err := ini.Load("conf/set/app.ini")
//	if err != nil {
//		Logger.Panic("Fail to parse 'conf/app.ini': %v", zap.Any("err", err))
//	}

//	if err = Cfg.Section("app").MapTo(AppSetting); err != nil {
//		Logger.Panic("Fail to map 'conf/app.ini': %v", zap.Any("err", err))
//	}
//
//	 5 MB check SetUp -> 1024 * 1024-> app.ini
//AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
//
//fmt.Println(AppSetting.ImageAllowExt) // v0.5.3 [] 没有成功获取; v0.5.4 [] 成功获取 app.ini的字段必须与结构体字段一致
//
//if err = Cfg.Section("server").MapTo(ServerSetting); err != nil {
//	Logger.Panic("Fail to map 'conf/app.ini': %v", zap.Any("err", err))
//}
//
//if err = Cfg.Section("mysql").MapTo(DatabaseSetting); err != nil {
//	Logger.Panic("Fail to map 'conf/app.ini': %v", zap.Any("err", err))
//}
//
//if err = Cfg.Section("redis").MapTo(RedisSetting); err != nil {
//	Logger.Panic("Fail to map 'conf/app.ini': %v", zap.Any("err", err))
//}
//
//ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
//ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
//
//Logger.Info("init conf success")
//}

//var (
//	Cfg          *ini.File
//	RunMode      string
//	HTTPPort     int
//	ReadTimeout  int
//	WriteTimeout int
//	PageSize     int
//	JwtSecret    string
//)

//func Init() {
//	var err error
//	Cfg, err = ini.Load("conf/app.ini")
//	if err != nil {
//		panic(err)
//	}
//	LoadBase()
//	LoadServer()
//	LoadApp()
//}
//
//func LoadBase() {
//	RunMode = Cfg.Section("").Key("run_mode").MustString("debug")
//}
//
//func LoadServer() {
//	sec, err := Cfg.GetSection("server")
//	if err != nil {
//		panic(err)
//	}
//	HTTPPort = sec.Key("http_port").MustInt(8088)
//	ReadTimeout = sec.Key("read_timeout").MustInt(60)
//	WriteTimeout = sec.Key("write_timeout").MustInt(60)
//}
//
//func LoadApp() {
//	sec, err := Cfg.GetSection("app")
//	if err != nil {
//		panic(err)
//	}
//	JwtSecret = sec.Key("jwt_secret").MustString("(❁´◡`❁)")
//	PageSize = sec.Key("page_size").MustInt(10)
//}
