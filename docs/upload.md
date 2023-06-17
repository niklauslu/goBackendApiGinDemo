### 文件上传

设置一下文件上传的配置

```go
// file: main.go
router := gin.Default()

// 设置上传
if err := setUploader(router); err != nil {
    logger.Fatalf("upload set err:%s", err.Error())
}
```

```go
// 上传设置
func setUploader(router *gin.Engine) error {

	if dir, err := os.Getwd(); err == nil {
		uploadPath = dir + "/uploads/"
	}

	if err := os.MkdirAll(uploadPath, 0777); err != nil {
		return err
	}

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", uploadPath)

	return nil
}
```

#### 上传接口示例（本地文件上传）

在全局配置中设置上传配置
```go
c.Set("upload_path", uploadPath)
c.Set("upload_host", os.Getenv("UPLOAD_HOST"))
```
接口示例
```go
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

```

