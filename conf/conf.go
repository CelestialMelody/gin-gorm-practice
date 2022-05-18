package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	*viper.Viper
}

var (
	GlobalConfig *Config
)

func init() {
	GlobalConfig = &Config{
		Viper: viper.New(),
	}
	// 初始化配置文件
	GlobalConfig.SetConfigName("app")
	GlobalConfig.SetConfigType("yaml")
	GlobalConfig.AddConfigPath(".")
	GlobalConfig.AddConfigPath("./conf/set")

	// 读取配置文件
	err := GlobalConfig.ReadInConfig()
	if err != nil {
		// WithField  分配一个新条目，并在其中添加了一个字段
		// WithError  将一个错误作为单个字段（使用ErrorKey中定义的键）添加到条目中
		logrus.WithField("config", "GlobalConfig").WithError(err).Panicf("unable to read global config")
	}

	// 监听配置文件变化
	GlobalConfig.WatchConfig()
	GlobalConfig.OnConfigChange(func(e fsnotify.Event) {
		err := GlobalConfig.ReadInConfig()
		if err != nil {
			logrus.WithField("config", "GlobalConfig").Info("config file update; change: ", e.Name)
		}
	})
}
