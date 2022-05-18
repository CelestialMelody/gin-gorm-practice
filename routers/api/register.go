package api

import (
	"gin-gorm-practice/models/blogAuth"
	"gin-gorm-practice/pkg/app"
	"gin-gorm-practice/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 用户注册
// @Summary 用户注册
// @Produce  json
// @Param name query string true "Name"
// @Param password query string true "Password"
// @Param email query string true "Email"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/register [post]
func Register(c *gin.Context) {
	type needValid struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}
	var need needValid
	appG := app.Gin{C: c}

	need.Username = c.Query("username")
	need.Password = c.Query("password")

	// 验证并处理参数
	if err := valid.Struct(need); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	if err := blogAuth.CheckAuth(need.Username, need.Password); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_USER, nil)
		// 创建用户
		err := blogAuth.Register(need.Username, need.Password)
		if err != nil {
			app.MarkError(err)
			appG.Response(http.StatusOK, e.ERROR_REGIEST_FAIL, nil)
			return
		}
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
