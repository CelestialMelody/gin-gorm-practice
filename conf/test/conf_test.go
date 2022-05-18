package T

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"testing"
	"time"
)

type Config struct {
	*viper.Viper
}

var (
	GlobalConfig *Config
	logger       = zap.NewExample()
)

func Init() {
	GlobalConfig = &Config{
		Viper: viper.New(),
	}
	// 初始化配置文件
	GlobalConfig.SetConfigName("app")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath("../set") // 搜索路径 yaml在当前目录

	// 读取配置文件
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		// WithField  分配一个新条目，并在其中添加了一个字段
		// WithError  将一个错误作为单个字段（使用ErrorKey中定义的键）添加到条目中
		logrus.WithField("config", "GlobalConfig").WithError(err).Panicf("unable to read global config")
		//logger.Panic("unable to read global config", zap.Error(err))
	}

	// 监听配置文件变化
	GlobalConfig.WatchConfig()
	GlobalConfig.OnConfigChange(func(e fsnotify.Event) {
		err := GlobalConfig.ReadInConfig()
		if err != nil {
			logrus.WithField("config", "GlobalConfig").Info("config file update; change: ", e.Name)
			//logger.Info("config file update;", zap.Any("change:", e.Name))
		}
	})
}

func TestInit(t *testing.T) {
	Init()
	fmt.Println(GlobalConfig.GetString("app.jwtSecret"))
	fmt.Println(GlobalConfig.GetString("mysql.mysql.username"))
	fmt.Println(GlobalConfig.GetString("mysql.redis.host"))

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

	var db DataBase

	_ = GlobalConfig.UnmarshalKey("db.mysql", &db)
	fmt.Printf("%+v\n", db)
}
