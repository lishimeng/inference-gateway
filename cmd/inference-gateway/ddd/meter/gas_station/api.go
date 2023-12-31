package gas_station

import (
	"fmt"
	"gitee.com/alex_li/inference-gateway/internal/etc"
	"gitee.com/alex_li/inference-gateway/internal/utils"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/pkg/errors"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
)

const (
	maxSize         = 50 * iris.MB
	uploadFormParam = "path"
)

// upload @[]path: 路径, @[]file: 文件
func upload(ctx iris.Context) (path, web string, err error) {

	ctx.SetMaxRequestBodySize(maxSize)

	uploadRoot, err := os.MkdirTemp("", "upload-*")
	if err != nil {
		log.Debug(errors.Wrapf(err, "create temp folder fail"))
		return
	}
	defer func() { // 清理缓存
		_ = os.RemoveAll(uploadRoot)
	}()

	uploaded, _, err := ctx.UploadFormFiles(uploadRoot)
	if err != nil {
		log.Debug(err)
		return
	}
	destDir := filepath.Join(etc.Config.FileSystem.Root)
	log.Debug("prepare to save files to:%s", destDir)

	err = os.MkdirAll(destDir, 0755)
	if err != nil {
		log.Debug(errors.Wrapf(err, "create dest dir fail:%s", destDir))
		return
	}

	var files []string
	for _, u := range uploaded { // 存储文件名
		files = append(files, u.Filename)
	}
	if len(uploaded) == 0 {
		// TODO 没文件
		return
	}
	var f = uploaded[0].Filename

	var destName string

	destName, err = copyFile(f, uploadRoot, destDir)
	if err != nil {
		log.Debug(errors.Wrapf(err, "save file fail:%s[%s]", destDir, destName))
		return
	}
	responsePath := filepath.Join("", destName)
	responsePath = filepath.ToSlash(responsePath)
	web, err = url.JoinPath(etc.Config.FileSystem.Domain, responsePath)
	if err != nil {
		log.Debug(errors.Wrapf(err, "save file fail:%s[%s]", destDir, destName))
		return
	}
	path = responsePath
	return
}

func copyFile(f, srcDir, destDir string) (fileName string, err error) {
	/// 名字处理
	ext := filepath.Ext(f)
	src := filepath.Join(srcDir, f)
	fileName, err = utils.FileDigest(src)
	if err != nil {
		return
	}
	fileName = fileName + ext
	///

	dest := filepath.Join(destDir, fileName)

	if err != nil {
		return
	}
	err = utils.CopyFile(src, dest)
	if err != nil {
		return
	}
	return
}

var testIndex = 0

func doInference(file string, inferenceKey string) (result map[string]any, err error) {

	testIndex = (testIndex + 1) % 5
	if testIndex == 0 {
		err = fmt.Errorf("inference err")
		return
	}
	result = make(map[string]any)
	result["totalPrice"] = 400.00
	result["quantity"] = 15.00
	result["unitPrice"] = 7.99
	return
}

type Response struct {
	app.Response
	Inference map[string]any `json:"inference,omitempty"`
	Image     string         `json:"image,omitempty"`
	ImageWeb  string         `json:"imageWeb,omitempty"`
}

var reg, _ = regexp.Compile("^[-\\w]+$")

func validParam(p string) bool {
	return reg.MatchString(p)
}

func inference(ctx iris.Context) {

	var err error
	var resp Response

	var wxId = ctx.GetHeader("WXID")
	var inferenceKey = ctx.GetHeader("Inference-key")
	if !validParam(wxId) || !validParam(inferenceKey) {

		resp.Code = 401
		resp.Message = "lost parameter"
		tool.ResponseJSON(ctx, resp)
		return
	}

	var imagePath, imageWeb string
	imagePath, imageWeb, err = upload(ctx)

	inferenceResult, err := doInference(imagePath, inferenceKey)
	if err != nil {
		resp.Code = 500
		resp.Message = "err"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Inference = inferenceResult
	resp.Image = imagePath
	resp.ImageWeb = imageWeb
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
