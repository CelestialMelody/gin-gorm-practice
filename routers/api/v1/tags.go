package v1

import (
	"gin-gorm-practice/conf"
	"gin-gorm-practice/pkg/app"
	"gin-gorm-practice/pkg/e"
	"gin-gorm-practice/pkg/util"
	"gin-gorm-practice/pkg/validate"
	"gin-gorm-practice/service/tagS"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// GetTagLists - 获取多个文章标签 GET
// @Summary GetTagLists
// @Produce json
// @Tags 标签
// @Description Get multiple blogArticle tags
// @Param name query string false "标签名称"
// @Param state query int false "状态"
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [get]
func GetTagLists(c *gin.Context) {
	type needValid struct {
		Name  string `validate:"max=100"`
		State int    `validate:"oneof=0 1"`
	}

	var need needValid
	appG := app.Gin{C: c}
	need.Name = c.Query("name")
	need.State = com.StrTo(c.Query("state")).MustInt()

	if err := validate.Struct(need); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tagServeice := tagS.Tag{
		Name:     need.Name,
		State:    need.State,
		PageNum:  util.GetPage(c),
		PageSize: conf.AppSetting.PageSize,
	}

	tagLists, err := tagServeice.GetAll()
	if err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
	}

	total, err := tagServeice.Count()
	if err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"lists": tagLists,
		"total": total,
	})
}

// AddTags - 添加多个文章标签 POST
// @Summary AddTags
// @Tags 标签
// @Description Add multiple blogArticle tags
// @Produce json
// @Param name query string true "标签名称"
// @Param state query int false "状态"
// @Param created_by query string true "创建人"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func AddTags(c *gin.Context) {
	type needValid struct {
		Name      string `validate:"required,max=100"`
		State     int    `validate:"oneof=0 1"`
		CreatedBy string `validate:"required,max=100"`
	}

	var need needValid
	appG := app.Gin{C: c}
	need.Name = c.Query("name")
	need.State = com.StrTo(c.Query("state")).MustInt()
	need.CreatedBy = c.Query("created_by")

	if err := validate.Struct(need); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tagService := tagS.Tag{
		Name:      need.Name,
		State:     need.State,
		CreatedBy: need.CreatedBy,
	}

	if err := tagService.ExistByName(); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}

	if err := tagService.Add(); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// EditTags - 编辑多个文章标签 PUT update
// @Summary EditTags
// @Tags 标签
// @Description Edit multiple blogArticle tags
// @Produce json
// @Param id path int true "ID"
// @Param name query string true "标签名称"
// @Param state query int false "状态"
// @Param modified_by query string true "修改人"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [put]
func EditTags(c *gin.Context) {
	type needValid struct {
		ID         int    `validate:"required,min=1"`
		Name       string `validate:"required,max=100"`
		State      int    `validate:"oneof=0 1"`
		ModifiedBy string `validate:"required,max=100"`
	}

	var need needValid
	need.ID = com.StrTo(c.Param("id")).MustInt()
	need.Name = c.Query("name")
	need.State = com.StrTo(c.Query("state")).MustInt()
	need.ModifiedBy = c.Query("modified_by")
	appG := app.Gin{C: c}

	if err := validate.Struct(need); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tagService := tagS.Tag{
		ID:         need.ID,
		Name:       need.Name,
		State:      need.State,
		ModifiedBy: need.ModifiedBy,
	}

	if err := tagService.ExistByID(); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}

	if err := tagService.Edit(); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

// DeleteTags - 删除多个文章标签
// @Summary DeleteTags
// @Tags 标签
// @Produce json
// @Param id path int true "ID"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags/{id} [delete]
func DeleteTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	appG := app.Gin{C: c}

	if err := validate.Var(id, "required,min=1"); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	tagService := tagS.Tag{ID: id}
	if err := tagService.ExistByID(); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG, nil)
		return
	}

	if err := tagService.Delete(); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_TAG_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
