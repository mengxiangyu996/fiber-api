package upload

import (
	"fiber-api/config"
	"fiber-api/pkg/storage"
	"encoding/base64"
	"errors"
	"io/ioutil"
	"mime/multipart"
	"strings"
)

// 上传文件
func File(fileHeader *multipart.FileHeader) (string, error) {

	file, _ := fileHeader.Open()
	fileByte, _ := ioutil.ReadAll(file)

	// 保存文件
	url, err := storage.Default().SetFileName(fileHeader.Filename).SetFileContent(fileByte).Save()

	return url, err
}

// 上传base64文件
func Base64(file, name string) (string, error) {

	fileByte, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return "", errors.New("base64文件转换失败")
	}

	// 保存文件
	url, err := storage.Default().SetFileName(name).SetFileContent(fileByte).Save()

	return url, err
}

// 重写url
func rewriteUrl(url string) string {

	if strings.HasPrefix(url, "http") {
		return url
	}

	return config.App.Domain + strings.ReplaceAll(url, "./web/app", "")
}
