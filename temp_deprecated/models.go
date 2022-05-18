package temp

//
//import (
//	"fmt"
//	"gin-gorm-practice/pkg/log"
//	"go.uber.org/zap"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"gorm.io/gorm/logger"
//	"gorm.io/gorm/schema"
//)
//
//var mysql *gorm.mysql
//
//type Model struct {
//	ID         int            `gorm:"primary_key" json:"id" validate:"min=1"`
//	CreatedON  string         `json:"created_on"` // 数据库时间改为varchar了
//	ModifiedON string         `json:"modified_on"`
//	DeleteOn   gorm.DeletedAt `json:"deleted_on" gorm:"column:deleted_on"`
//}
//
//type BeforeDB interface {
//	BeforeCreate(mysql *gorm.mysql) error
//	BeforeUpdate(mysql *gorm.mysql) error
//}
//
//func SetUp() {
//	mysql, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
//var err error
//mysql, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
//	DatabaseSetting.User,
//	DatabaseSetting.Password,
//	DatabaseSetting.Host,
//	DatabaseSetting.Port,
//	DatabaseSetting.Name)), &gorm.Config{
//	NamingStrategy: schema.NamingStrategy{
//		TablePrefix:   DatabaseSetting.TablePrefix, // 数据库表前缀
//		SingularTable: true,                        // 使用单数表名
//	},
//	Logger: logger.Default, // 日志级别
//})
//if err != nil {
//	log.Logger.Error("数据库连接失败", zap.Any("err", err))
//}
//mysqlDB, _ := mysql.mysql()
//SetMaxIdleConns 设置空闲连接池中连接的最大数量
//mysqlDB.SetMaxIdleConns(10)
//SetMaxOpenConns 设置打开数据库连接的最大数量。
//mysqlDB.SetMaxOpenConns(100)
//SetConnMaxLifetime 设置连接的最大可复用时间，超过时间的连接会被关闭。
//mysqlDB.SetConnMaxLifetime(100)
//
//return mysql
//}

//func InitDatabase() *gorm.mysql {
//	var (
//		err                                       error
//		dbName, user, password, host, tablePrefix string
//		port                                      int
//	)
//	sec, err := set.Cfg.GetSection("mysql") // app.ini
//	if err != nil {
//		log.Fatal(2, " Fail to get section 'database': %v", err)
//	}
//
//	//dbType = sec.Key("TYPE").String()
//	// app.ini
//	dbName = sec.Key("dbname").String()
//	user = sec.Key("user").String()
//	password = sec.Key("password").String()
//	host = sec.Key("host").String()
//	port, _ = sec.Key("port").Int()
//	tablePrefix = sec.Key("table_prefix").String() //数据库表前缀
//
//	fmt.Println(dbName, user, password, host, port, tablePrefix)
//
//	// gorm v2 用法
//	mysql, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
//		user,
//		password,
//		host,
//		port,
//		dbName)), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			TablePrefix:   tablePrefix, // 数据库表前缀
//			SingularTable: true,        // 使用单数表名
//		},
//		// Gorm v2
//		//Logger: logger.Default.LogMode(logger.Silent), // 日志级别
//		Logger: logger.Default, // 日志级别
//	})
//	/*
//		Gorm V1 有内置的日志记录器支持，默认情况下，它会打印发生的错误
//		// 启用Logger，显示详细日志
//		mysql.LogMode(true)
//	*/
//	if err != nil {
//		//log.Fatal(2, " Fail to connect database: %v", err)
//		logrus.Fatal("Fail to connect database: %v", err)
//		//logrus.Println(err)
//		//return
//	}
//	logrus.Infof("Connected mysql successfully, dbname: %s", dbName)
//	// Gorm 2.0 用法
//	mysqlDB, err := mysql.mysql()
//	if err != nil {
//		//log.Println(err)
//		logrus.Println(err)
//	}
//	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
//	mysqlDB.SetMaxIdleConns(10)
//	// SetMaxOpenConns 设置打开数据库连接的最大数量。
//	mysqlDB.SetMaxOpenConns(100)
//	// SetConnMaxLifetime 设置连接的最大可复用时间，超过时间的连接会被关闭。
//	mysqlDB.SetConnMaxLifetime(100)
//
//	// 注册回调函数
//	//err = mysql.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
//	return mysql
//}
