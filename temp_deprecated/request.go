package temp

import (
	"gin-gorm-practice/pkg/log"
	"github.com/beego/beego/v2/core/validation"
	"go.uber.org/zap"
)

func MarkError(err interface{}) {
	switch err.(type) {
	case error:
		e := err.(error)
		log.Logger.Error("message:", zap.Error(e))
	case *validation.Error: // 之前使用过beego的validation 版本1.0开始不使用了
		e := err.(*validation.Error)
		log.Logger.Error(e.Key, zap.Any(";massage:", e.Value))
	}
}
