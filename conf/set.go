package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

type App struct {
	Port            string
	Host            string
	JwtSecret       string
	RuntimeRootPath string
	PageSize        int
	MaxHeaderBytes  int
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	Release         bool
	RunMode         string
}

type DataBase struct {
	Address     string
	User        string
	Password    string
	Database    string
	MaxIdle     int
	MaxOpen     int
	MaxLifetime time.Duration
	TablePrefix string
}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type Image struct {
	ImagePrefixUrl string
	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExt  []string
}

type Log struct {
	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var (
	AppSetting      App
	DataBaseSetting DataBase
	RedisSetting    Redis
	ImageSetting    Image
	LogSetting      Log
)

func Init() {
	// 初始化配置
	var err error

	if err = GlobalConfig.UnmarshalKey("app", &AppSetting); err != nil {
		logrus.Panicf("解析配置文件失败 app: %v", err)
	}

	if err = GlobalConfig.UnmarshalKey("db.mysql", &DataBaseSetting); err != nil {
		logrus.Panicf("解析配置文件失败 database: %v", err)
	}

	if err = GlobalConfig.UnmarshalKey("db.redis", &RedisSetting); err != nil {
		logrus.Panicf("解析配置文件失败 redis: %v", err)
	}

	if err = GlobalConfig.UnmarshalKey("image", &ImageSetting); err != nil {
		logrus.Panicf("解析配置文件失败 image: %v", err)
	}

	if err = GlobalConfig.UnmarshalKey("log", &LogSetting); err != nil {
		logrus.Panicf("解析配置文件失败 log: %v", err)
	}

	AppSetting.ReadTimeout = AppSetting.ReadTimeout * time.Second
	AppSetting.WriteTimeout = AppSetting.WriteTimeout * time.Second
	AppSetting.MaxHeaderBytes = AppSetting.MaxHeaderBytes << 20
	ImageSetting.ImageMaxSize = ImageSetting.ImageMaxSize * 1024 * 1024

	fmt.Printf("%+v\n", DataBaseSetting)
}
