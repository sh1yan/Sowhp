package concert

import (
	"Sowhp/concert/logger"
	"fmt"
	"os"
)

// MkdirResport 创建报告存放目录
func MkdirResport() {
	// 获取当前目录地址
	dir, err := os.Getwd()
	if err != nil {
		// panic(err)
		logger.DebugError(err)
	}
	logger.Debug(fmt.Sprintf("当前路径:%s", dir))
	Dir_mk(dir + "/result/")
	logger.Debug("当前目录下创建 /result/ 目录成功！")

}

// GetPath 获取当前路径位置
func GetPath() string {
	// 获取当前目录地址
	dir, err := os.Getwd()
	if err != nil {
		// panic(err)
		logger.DebugError(err)
	}
	return dir
}

// dir_mk 判断目录是否存在，若不存在则进行创建
func Dir_mk(path string) {
	// 判断目录是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0777)
		if err != nil {
			// panic(err)
			logger.DebugError(err)
		}
		logger.Debug(fmt.Sprintf("当前以创建好该目录：%s", path))
		return
	}
}
