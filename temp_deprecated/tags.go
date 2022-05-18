package temp

//
//import (
//	"gin-gorm-practice/conf"
//	"gin-gorm-practice/models/blogTag"
//	"gin-gorm-practice/pkg/app"
//	"gin-gorm-practice/pkg/e"
//	"gin-gorm-practice/pkg/log"
//	"gin-gorm-practice/pkg/util"
//	"gin-gorm-practice/pkg/validate"
//	"gin-gorm-practice/service/tagS"
//	"github.com/gin-gonic/gin"
//	"github.com/unknwon/com"
//	"go.uber.org/zap"
//	"net/http"
//)
//
//// GetTagLists - 获取多个文章标签 GET
//// @Summary GetTagLists
//// @Produce json
//// @Tags 标签
//// @Description Get multiple blogArticle tags
//// @Param name query string false "标签名称"
//// @Param state query int false "状态"
//// @Param page query int false "页码"
//// @Param page_size query int false "每页数量"
//// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//// @Router /api/v1/tags [get]
//func GetTagLists(c *gin.Context) {
//	type needValid struct {
//		Name  string `validate:"max=100"`
//		State int    `validate:"oneof=0 1"`
//	}
//
//	var need needValid
//	appG := app.Gin{C: c}
//	need.Name = c.Query("name")
//	need.State = com.StrTo(c.Query("state")).MustInt()
//
//	if err := validate.Struct(need); err != nil {
//		app.MarkError(err)
//		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
//		return
//	}
//
//	tagServeice := tagS.Tag{
//		Name:     need.Name,
//		State:    need.State,
//		PageNum:  util.GetPage(c),
//		PageSize: conf.AppSetting.PageSize,
//	}
//
//	tagLists, err := tagServeice.GetAll()
//	if err != nil {
//		app.MarkError(err)
//		appG.Response(http.StatusInternalServerError, e.ERROR_GET_TAGS_FAIL, nil)
//	}
//
//	total, err := tagServeice.Count()
//	if err != nil {
//		app.MarkError(err)
//		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_TAG_FAIL, nil)
//	}
//
//	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
//		"lists": tagLists,
//		"total": total,
//	})
//}
//
////AddTags - 添加多个文章标签 POST
////@Summary AddTags
////@Tags 标签
////@Description Add multiple blogArticle tags
////@Produce json
//// @Param name query string true "标签名称"
//// @Param state query int false "状态"
//// @Param created_by query string true "创建人"
//// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//// @Router /api/v1/tags [post]
//func AddTags(c *gin.Context) {
//	name := c.Query("name")
//	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
//	createdBy := c.Query("created_by")
//
//	valid := validation.Validation{}
//	valid.Required(name, "name").Message("名称不能为空")
//	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
//	valid.Required(createdBy, "created_by").Message("创建人不能为空")
//	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
//	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
//
//	code := e.INVALID_PARAMS
//
//	type needValid struct {
//		Name      string `validate:"required,max=100"`
//		State     int    `validate:"oneof=0 1"`
//		CreatedBy string `validate:"required,max=100"`
//	}
//
//	if !valid.HasErrors() {
//		zap.L().Debug("AddTags: ", zap.Any("validate", valid.HasErrors()), zap.Any("validate.Errors", valid.Errors))
//		if !blogTag.ExistTagByName(name) { // 判断是否存在
//			if blogTag.AddTag(name, state, createdBy) { // 添加
//				code = e.SUCCESS
//			} else {
//				//log.Info("添加多个文章标签失败")
//				log.Logger.Info("添加多个文章标签失败")
//				code = e.ERROR_ADD_TAG // 添加失败
//			}
//		} else {
//			code = e.ERROR_EXIST_TAG           // 已存在
//			for _, err := range valid.Errors { // demo 测试自己的日志
//				//log.Info(err.Key, err.Message)
//				log.Logger.Info("err", zap.Any("err key", err.Key), zap.Any("err value", err.Value))
//			}
//		}
//	} else {
//		for _, err := range valid.Errors { // demo 测试自己的日志
//			//log.Info(err.Key, err.Message)
//			log.Logger.Error(
//				"AddTags: ",
//				zap.Any("err", err.Key),
//				zap.Any("err", err.Message),
//			)
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
//
//// EditTags - 编辑多个文章标签 PUT update
//// @Summary EditTags
//// @Tags 标签
//// @Description Edit multiple blogArticle tags
//// @Produce json
//// @Param id path int true "ID"
//// @Param name query string true "标签名称"
//// @Param state query int false "状态"
//// @Param modified_by query string true "修改人"
//// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//// @Router /api/v1/tags/{id} [put]
//func EditTags(c *gin.Context) {
//	// 获取参数
//	id := com.StrTo(c.Param("id")).MustInt()
//	name := c.Query("name")
//	modifiedBy := c.Query("modified_by")
//
//	// 验证参数
//	valid := validation.Validation{}
//	state := -1
//
//	if arg := c.Query("state"); arg != "" {
//		state = com.StrTo(arg).MustInt()
//		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
//	}
//
//	// 验证参数
//	valid.Required(id, "id").Message("ID不能为空")
//	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
//	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
//	valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
//
//	code := e.INVALID_PARAMS
//	if !valid.HasErrors() {
//		if blogTag.ExistTagByID(id) {
//			code = e.SUCCESS
//			// 与AddTag 写法不同 也可以类似于AddTag那样写
//			data := make(map[string]interface{})
//			data["modified_by"] = modifiedBy
//			if name != "" {
//				data["name"] = name
//			}
//			if state != -1 {
//				data["state"] = state
//			}
//			blogTag.EditTag(id, data) //
//		} else {
//			code = e.ERROR_NOT_EXIST_TAG
//			log.Logger.Error("EditTags: ", zap.Any("err", "标签不存在"))
//		}
//	} else {
//		for _, err := range valid.Errors { // demo 测试自己的日志
//			//log.Info(err.Key, err.Message)
//			log.Logger.Error(
//				"EditTags: ",
//				zap.Any("err", err.Key),
//				zap.Any("err", err.Message),
//			)
//		}
//	}
//
//	// 返回结果
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
//
//// DeleteTags - 删除多个文章标签
//// @Summary DeleteTags
//// @Tags 标签
//// @Produce json
//// @Param id path int true "ID"
//// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
//// @Router /api/v1/tags/{id} [delete]
//func DeleteTags(c *gin.Context) {
//	id := com.StrTo(c.Param("id")).MustInt()
//	log.Logger.Debug("id", zap.Any("id", id))
//
//	valid := validation.Validation{}
//	valid.Required(id, "id").Message("ID不能为空")
//	valid.Min(id, 1, "id").Message("ID必须大于0")
//
//	//binding.Validator = new(blogTag.Tag)
//	code := e.INVALID_PARAMS
//	if !valid.HasErrors() {
//		if blogTag.ExistTagByID(id) {
//			code = e.SUCCESS
//			blogTag.DeleteTag(id)
//		} else {
//			code = e.ERROR_NOT_EXIST_ARTICLE
//			log.Logger.Error("err", zap.Any("err", "文章不存在"))
//		}
//	} else {
//		for _, err := range valid.Errors { // demo 测试自己的日志
//			//log.Info(err.Key, err.Message)
//			log.Logger.Error(
//				"DeleteTags: ",
//				zap.Any("err", err.Key),
//				zap.Any("err", err.Message),
//			)
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": make(map[string]string),
//	})
//}
//
//// 打印不出来 测试
////logger := zap.NewExample()
////logger.Info("AddTags: ", zap.Any("err", err.Key), zap.Any("err", err.Message))
////zap.S().Error("AddTags: ", zap.Any("err", err.Key), zap.Any("err", err.Message))
