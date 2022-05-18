package tagS

import (
	"gin-gorm-practice/models/blogTag"
	"gin-gorm-practice/pkg/app"
	"gin-gorm-practice/pkg/redis"
	"gin-gorm-practice/service/cache"
	jsoniter "github.com/json-iterator/go"
)

type Tag struct {
	ID         int
	Name       string
	CreatedBy  string
	ModifiedBy string
	State      int

	PageNum  int
	PageSize int
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func (t *Tag) ExistByName() error {
	return blogTag.ExistTagByName(t.Name)
}

func (t *Tag) ExistByID() error {
	return blogTag.ExistTagByID(t.ID)
}

func (t *Tag) Add() error {
	return blogTag.AddTag(t.Name, t.State, t.CreatedBy)
}

func (t *Tag) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = t.ModifiedBy
	data["name"] = t.Name
	if t.State >= 0 {
		data["state"] = t.State
	}

	return blogTag.EditTag(t.ID, data)
}

func (t *Tag) Delete() error {
	return blogTag.DeleteTag(t.ID)
}

func (t *Tag) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})

	if t.Name != "" {
		maps["name"] = t.Name
	}

	if t.State >= 0 {
		maps["state"] = t.State
	}

	return maps
}

func (t *Tag) Count() (int64, error) {
	return blogTag.GetTagCount(t.getMaps())
}

func (t *Tag) GetAll() ([]blogTag.Tag, error) {
	tagCache := cache.Tag{
		State:    t.State,
		PageNum:  t.PageNum,
		PageSize: t.PageSize,
	}
	key := tagCache.GetKey()
	if redis.Exists(key) {
		data, err := redis.Get(key)
		if err != nil {
			app.MarkError(err)
		} else {
			var tags []blogTag.Tag
			_ = json.Unmarshal(data, &tags)
			return tags, nil
		}
	}

	tags, err := blogTag.GetTags(t.PageNum, t.PageSize, t.getMaps())
	if err != nil {
		return nil, err
	}

	if err := redis.Set(key, tags, 3600); err != nil {
		app.MarkError(err)
	}

	return tags, nil
}

//// Export 导出excel
//func (t *Tag) Export() (string, error) {
//	tags, err := t.GetAll()
//	if err != nil {
//		return "", err
//	}
//
//	xlsFile := xlsx.NewFile()
//}
