package blogArticle

import (
	"fmt"
	"gin-gorm-practice/models"
	"gin-gorm-practice/models/blogTag"
	"gin-gorm-practice/pkg/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

var instanceDB *models.DBList

// 接口可以返回bool 方便处理 返回err 感觉处理时太丑了 err1 err2...; 有必要的接口返回err; 看看别的项目 学习一下
// 是我傻了 返回的err是局部变量 不冲突

// Article v0.2.4之前为 BlogArticle 但在配置时写了前缀 故改为 Article
type Article struct {
	models.Module
	TagID         int         `json:"tag_id" gorm:"index" validate:"min=1"`
	Tag           blogTag.Tag `json:"tag"`
	Title         string      `json:"title" validate:"min=1,max=100"`
	Desc          string      `json:"desc" validate:"min=1,max=100"`
	Content       string      `json:"content" validate:"min=1"`
	CreatedBy     string      `json:"created_by" validate:"min=1,max=100"`
	ModifiedBy    string      `json:"modified_by" validate:"min=1,max=100"`
	CoverImageUrl string      `json:"cover_image_url" validate:"min=1,max=255"` // v0.5 增加封面图片c
	//DeletedOn     string      `json:"deleted_on" validate:"min=1,max=100"`
	// 做软删除
	State int `json:"state" validate:"oneof=0 1"`
}

func init() {
	instanceDB = models.InitDB()
	if err := instanceDB.MysqlDB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Article{}); err != nil {
		logging.LoggoZap.Error(fmt.Sprintf("MysqlDB.AutoMigrate error: %v", err))
	}
}

// BeforeCreate 建议抽象为接口
func (article *Article) BeforeCreate(db *gorm.DB) error {
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	db.Statement.SetColumn("created_on", fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second))
	return nil
}

func (article *Article) BeforeUpdate(db *gorm.DB) error { // 大写
	year := time.Now().Year()
	month := time.Now().Month()
	day := time.Now().Day()
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	db.Statement.SetColumn("modified_on", fmt.Sprintf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second))
	return nil
}

// ExistArticleByID 根据ID查询文章是否存在; id 与 tag_id?
func ExistArticleByID(id int) error {
	var article Article
	err := instanceDB.MysqlDB.Select("id").Where("id = ?", id).First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	if article.TagID > 0 {
		return nil
	}
	return nil
}

// GetArticleTotalCount 查询文章总数
func GetArticleTotalCount(maps interface{}) (count int64) {
	instanceDB.MysqlDB.Model(&Article{}).Where(maps).Count(&count)
	return
}

// GetArticles 获取文章列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	// DB.Preload 查询关联表; blog_blog_article; 问题: 为什么大写Tag?
	instanceDB.MysqlDB.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)
	return
}

// GetArticle 获取文章
func GetArticle(id int) (article Article) {
	instanceDB.MysqlDB.Where("id = ?", id).First(&article)
	// DB.Association 关联查询
	err := instanceDB.MysqlDB.Model(&article).Association("tag").Find(&article.Tag)
	if err != nil {
		logrus.Debugln("Can't Find Article", err)
		return Article{}
	}
	return
}

func AddArticle(data map[string]interface{}) error {
	instanceDB.MysqlDB.Create(&Article{
		//map[string]interface{}.(type) 接口类型断言
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})
	return nil
}

// EditArticle 编辑文章; 只会返回true 不太好
func EditArticle(id int, data interface{}) bool {
	instanceDB.MysqlDB.Model(&Article{}).Where("id = ?", id).Updates(data)
	return true
}

// DeleteArticle 删除文章 只会返回true 不太好
func DeleteArticle(id int) bool {
	instanceDB.MysqlDB.Where("id = ?", id).Delete(&Article{})
	return true
}

func CleanAllArticle() bool {
	instanceDB.MysqlDB.Unscoped().Where("deleted_on != ?", 0).Delete(&Article{})
	return true
}

//func init() {
//	instanceDB = models.InitDB()
//logger := zap.NewExample().Sugar()
// 创建表 Cannot add foreign key constraint; 放弃使用auto migrate
// 解决 循环依赖
//if err := instanceDB.MysqlDB.AutoMigrate(&Article{}); err != nil {
//	logger.Error("Article AutoMigrate error", zap.Error(err))
//}
//instanceDB.MysqlDB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
//logger.Infof("Article AutoMigrate success")
//}

// Article 通过AutoMigrate()创建表 db.set() 如果要用到外键 貌似不好创建; 放弃
//type Article struct { // gorm 字段设置; comment 注释 add comment for field when migration
//	models.Module
//
//	// Cannot add foreign key constraint;
//	TagID uint        `json:"tag_id" gorm:"column:tag_id;type:int(10) unsigned;not null;default:0;comment:'标签ID'" binding:"required"`
//	Tag   blogTag.Tag `json:"tag"`
//	Title string      `json:"title" gorm:"column:title;type:varchar(100);not null;default:'';comment:'文章标题'" binding:"required"`
//
//	Desc       string `json:"desc" gorm:"column:desc;type:varchar(255);not null;default:'';comment:'简述'" binding:"required"`
//	Content    string `json:"content" gorm:"column:content;type:text;comment:'内容'" binding:"required"`
//	CreatedBy  string `json:"created_by" gorm:"column:created_by;type:varchar(100);not null;default:'';comment:'创建人'" binding:"required"`
//	ModifiedBy string `json:"modified_by" gorm:"column:modified_by;type:varchar(255);not null;default:'';comment:'修改人'" binding:"required"`
//	// 实际上 这是硬删除 没法用
//	DeletedOn string `json:"deleted_on" gorm:"column:deleted_on;type:varchar(100);not null;default:'';comment:'删除时间'" binding:"required"`
//	State     int    `json:"state" gorm:"column:state;type:tinyint(3);not null;default:1;comment:'状态 0为禁用1为启用'" binding:"required"`
//}

//CREATE TABLE `blog_article` (
//`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
//`tag_id` int(10) unsigned DEFAULT '0' COMMENT '标签ID',
//`title` varchar(100) DEFAULT '' COMMENT '文章标题',
//`desc` varchar(255) DEFAULT '' COMMENT '简述',
//`content` text,
//`created_on` varchar(100) DEFAULT '' COMMENT '创建时间',
//`created_by` varchar(100) DEFAULT '' COMMENT '创建人',
//`modified_on` varchar(100) DEFAULT '' COMMENT '修改时间',
//`modified_by` varchar(100) DEFAULT '' COMMENT '修改人',
//`deleted_on` varchar(100) DEFAULT '' COMMENT '删除时间',
//`state` tinyint(3) unsigned DEFAULT '1' COMMENT '状态 0为禁用1为启用',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章管理';
