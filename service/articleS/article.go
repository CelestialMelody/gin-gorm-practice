package articleS

import (
	"gin-gorm-practice/models/blogArticle"
	"gin-gorm-practice/pkg/app"
	"gin-gorm-practice/pkg/redis"
	"gin-gorm-practice/service/cache"
	jsoniter "github.com/json-iterator/go"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}

	err := blogArticle.AddArticle(article)
	return err

}

func (a *Article) Edit() error {
	return blogArticle.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

// 先从缓存中取; 仅对redis的error进行处理
// 缓存中没有再从数据库中取

func (a *Article) Get() (*blogArticle.Article, error) {
	articleCache := cache.Article{ID: a.ID}
	key := articleCache.GetArticleKey()
	if redis.Exists(key) {
		data, err := redis.Get(key)
		if err != nil {
			app.MarkError(err)
		} else {
			var article *blogArticle.Article
			_ = json.Unmarshal(data, &article)
			return article, nil
		}
	}

	article, err := blogArticle.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}

	if err := redis.Set(key, article, 3600); err != nil {
		app.MarkError(err)
	}
	return article, nil
}

func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}
	return maps
}

func (a *Article) GetAll() ([]*blogArticle.Article, error) {
	articlesCache := cache.Article{
		TagID:    a.TagID,
		State:    a.State,
		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}

	key := articlesCache.GetArticlesKey()
	if redis.Exists(key) {
		data, err := redis.Get(key)
		if err != nil {
			app.MarkError(err)
		} else {
			var articles []*blogArticle.Article
			_ = json.Unmarshal(data, &articles)
			return articles, nil
		}
	}

	articles, err := blogArticle.GetArticleLists(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	if err := redis.Set(key, articles, 3600); err != nil {
		app.MarkError(err)
	}
	return articles, nil
}

func (a *Article) Delete() error {
	return blogArticle.DeleteArticle(a.ID)
}

func (a *Article) ExistByID() error {
	return blogArticle.ExistArticleByID(a.ID)
}

func (a *Article) Count() (int64, error) {
	return blogArticle.GetArticleTotalCount(a.getMaps())
}
