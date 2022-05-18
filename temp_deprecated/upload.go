package temp

import (
	"fmt"
	"gin-gorm-practice/pkg/app"
	"gin-gorm-practice/pkg/e"
	"gin-gorm-practice/pkg/log"
	"gin-gorm-practice/pkg/upload"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

// UploadImage upload image
// @Summary Upload image
// @Tags 图片上传
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "{"code":0,"message":"ok","data":{}}"
// @Router /upload [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}

	file, fileHeader, err := c.Request.FormFile("image") // 获取上传文件
	if err != nil {
		app.MarkError(err)
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_IMAGE_FAIL, nil)
		return
	}

	if fileHeader == nil {
		err = fmt.Errorf("image is nil")
		app.MarkError(err)
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_IMAGE_FAIL, nil)
		return
	}

	imageName := upload.GetImageName(fileHeader.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImageSavePath()
	src := fullPath + imageName

	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		if !upload.CheckImageExt(imageName) {
			err = fmt.Errorf("image format is invalid")
			app.MarkError(err)
			appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		}
		if !upload.CheckImageSize(file) {
			err = fmt.Errorf("image size is invalid")
			app.MarkError(err)
			appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_SIZE, nil)
		}
		return
	}

	if err := upload.CheckImage(fullPath); err != nil {
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	if err := c.SaveUploadedFile(fileHeader, src); err != nil { // 保存文件
		app.MarkError(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_IMAGE_FAIL, nil)
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}

// UploadImages UploadFile upload file
// @Summary Upload file
// @Tags 文件上传
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "{"code":0,"message":"ok","data":{}}"
// @Router /upload/file [post]
func UploadImages(c *gin.Context) {
	code := e.SUCCESS
	data := make(map[string]interface{})

	//forms := c.Request.MultipartForm
	//files, ok := forms.File["images"]
	forms, _ := c.MultipartForm()
	files, ok := forms.File["images"]
	if !ok {
		code = e.INVALID_PARAMS
		err := fmt.Errorf("images is nil")
		log.Logger.Error("Invalid params", zap.Any("err", err))
	}
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			code = e.ERROR
			err = fmt.Errorf("failed to get image")
			log.Logger.Error("Failed to get image", zap.Any("err", err))

			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
		}
		imageName := upload.GetImageName(files[i].Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImageSavePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
			if !upload.CheckImageExt(imageName) {
				err = fmt.Errorf("image format is invalid")
				log.Logger.Error("Image format is invalid", zap.Any("err", err))
			}
			if !upload.CheckImageSize(file) {
				err = fmt.Errorf("image size is invalid")
				log.Logger.Error("Image size is invalid", zap.Any("err", err))
			}
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
				log.Logger.Error("Check image failed", zap.Any("err", err))
			} else if err := c.SaveUploadedFile(files[i], src); err != nil { // 保存文件
				code = e.ERROR_UPLOAD_IMAGE_FAIL
				log.Logger.Error("Save image failed", zap.Any("err", err))
			} else {
				data["image_url_"+strconv.Itoa(i)] = upload.GetImageFullUrl(imageName)
				data["image_save_url_"+strconv.Itoa(i)] = savePath + imageName

				c.JSON(http.StatusOK, gin.H{
					"code": code,
					"msg":  e.GetMsg(code),
					"data": data,
				})
			}
		}
	}
}
