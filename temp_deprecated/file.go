package temp

import (
	"fmt"
	"gin-gorm-practice/conf"
	"gin-gorm-practice/pkg/file"
	"go.uber.org/zap"
	"os"
	"time"
)

var logger = zap.NewExample()

func getLogFilePath() string {
	dir, _ := os.Getwd()
	return fmt.Sprintf("%s/%s%s",
		dir,
		conf.AppSetting.RuntimeRootPath,
		conf.LogSetting.LogSavePath)
}

func getLogFileName() string {
	return fmt.Sprintf("%s%s.%s",
		conf.LogSetting.LogSaveName,
		time.Now().Format(conf.LogSetting.TimeFormat),
		conf.LogSetting.LogFileExt,
	)
}

// getLogFileName get the save name of the log file
func getLogFileFullPath() string {
	preFixPath := getLogFilePath()
	suffixPath := getLogFileName()
	return fmt.Sprintf("%s%s", preFixPath, suffixPath)
}

func makeDir() error {
	if err := os.MkdirAll(getLogFilePath(), os.ModePerm); err != nil {
		logger.Debug("创建日志目录失败", zap.Error(err))
		return err
	}
	logger.Debug("创建日志目录成功", zap.String("dir", getLogFilePath()))
	return nil
}

func openLogFile(fileName, filePath string) (*os.File, error) {
	path := getLogFilePath()
	if perm := file.CheckPermission(path); perm == true {
		logger.Debug("日志目录权限不够", zap.String("path", path))
		return nil, fmt.Errorf("file.CheckPermission Permission denied , path: %s", path)
	}
	if err := file.IsNotExistMkDir(path); err != nil {
		logger.Debug("创建日志目录失败", zap.Error(err))
		return nil, fmt.Errorf("file.IsNotExistMkDir failed , path : %s", path)
	}
	f, err := file.Open(path+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Debug("打开日志文件失败", zap.Error(err))
		return nil, fmt.Errorf("fail to OpenFile :%v", err)
	}
	return f, nil
}
