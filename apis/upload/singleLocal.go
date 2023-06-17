package apis_upload

import (
	"fmt"
	"net/http"
	"niklauslu/goBackendApiGinDemo/lib"
	"niklauslu/goBackendApiGinDemo/utils"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var logger = lib.Logger("upload")

func SingleLocalUpload(c *gin.Context) {
	f, err := c.FormFile("file")
	if err != nil {
		logger.Errorf("%s|SingleLocalUpload|err:%+v", c.GetHeader("Request-Id"), err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("upload fail:%s", err.Error()))
		return
	}
	uploadPath := c.GetString("upload_path")

	now := time.Now()
	date := now.Format("2006-01-02")
	savePath := path.Join(uploadPath, date)
	exist := pathExists(savePath)
	if exist == false {
		os.MkdirAll(savePath, 0777)
	}

	dstFilename := utils.GenerateUUID() + "." + getFileExt(f.Filename)
	dst := path.Join(savePath, dstFilename)
	if err = c.SaveUploadedFile(f, dst); err != nil {
		logger.Errorf("%s|SingleLocalUpload|err:%+v", c.GetHeader("Request-Id"), err)
		c.JSON(http.StatusBadRequest, fmt.Sprintf("upload fail:%s", err.Error()))
		return
	}

	showPath := path.Join("/uploads", date, dstFilename)
	URL := c.GetString("upload_host") + showPath

	logger.Infof("%s|SingleLocalUpload|URL:%+v", c.GetHeader("Request-Id"), URL)
	c.JSON(http.StatusOK, gin.H{
		"path": showPath,
		"url":  URL,
	})
	return
}

// 判断所给路径文件/文件夹是否存在1
func pathExists(filepath string) bool {
	_, err := os.Stat(filepath) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 获取文件后缀
func getFileExt(filePath string) string {
	fileArr := strings.Split(filePath, ".")
	ext := fileArr[len(fileArr)-1]
	return ext
}
