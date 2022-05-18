package mysql

import (
	"fmt"
	"gin-gorm-practice/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Model struct {
	gorm.Model
}

var (
	DB *gorm.DB
)

func Init() (err error) {
	DB, err = CreateDB(struct {
		Addr string
		User string
		Pass string
		DB   string
	}{
		Addr: conf.DataBaseSetting.Address,
		User: conf.DataBaseSetting.User,
		Pass: conf.DataBaseSetting.Password,
		DB:   conf.DataBaseSetting.Database,
		//Addr: conf.GlobalConfig.GetString("db.mysql.address"),
		//User: conf.GlobalConfig.GetString("db.mysql.user"),
		//Pass: conf.GlobalConfig.GetString("db.mysql.password"),
		//DB:   conf.GlobalConfig.GetString("db.mysql.database"),
	})

	if err != nil {
		logrus.Panic("connect mysql error: ", err.Error())
		return err
	}

	logrus.Infof("Connected mysql success")

	db, _ := DB.DB()
	db.SetMaxIdleConns(conf.DataBaseSetting.MaxIdle)        // 设置最大空闲连接数
	db.SetMaxOpenConns(conf.DataBaseSetting.MaxOpen)        // 设置最大连接数
	db.SetConnMaxLifetime(conf.DataBaseSetting.MaxLifetime) // 设置连接最大存活时间

	return nil
}

func CreateDSN(dbInfo struct {
	Addr string
	User string
	Pass string
	DB   string
}) string {
	//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbInfo.User, dbInfo.Pass, dbInfo.Addr, dbInfo.DB)
}

func CreateDB(dbInfo struct {
	Addr string
	User string
	Pass string
	DB   string
}) (*gorm.DB, error) {
	cfg := struct {
		Addr string
		User string
		Pass string
		DB   string
	}{
		Addr: dbInfo.Addr,
		User: dbInfo.User,
		Pass: dbInfo.Pass,
		DB:   dbInfo.DB,
	}
	tablePrefix := conf.GlobalConfig.GetString("mysql.mysql.tablePrefix")
	DB, err := gorm.Open(mysql.Open(CreateDSN(cfg)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, // 数据库表前缀
			SingularTable: true,        // 使用单数表名
		},
		PrepareStmt: true,           // 预处理语句
		Logger:      logger.Default, // 日志级别
	})
	return DB, err
}
