package utils

import (
	"os"
)

// IsFileOrDirExists 判断文件或文件夹是否存在
func IsFileOrDirExists(src string) bool {
	_, err := os.Stat(src)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}

	return true
}

// IsFile 是否是文件
func IsFile(src string) bool {
	f, err := os.Stat(src)
	if err != nil {
		return false
	}

	return !f.IsDir()
}

// IsDir 是否是目录
func IsDir(src string) bool {
	f, err := os.Stat(src)
	if err != nil {
		return false
	}

	return f.IsDir()
}
