package api

import (
	"EvelyApi/app"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/goadesign/goa"
	"io"
	"os"
	"strings"
	"time"
)

// FilesController implements the files resource.
type FilesController struct {
	*goa.Controller
	// db *model.FileDB
}

// NewFilesController creates a files controller.
func NewFilesController(service *goa.Service) *FilesController {
	return &FilesController{Controller: service.NewController("FilesController")}
}

// Upload runs the upload action.
func (c *FilesController) Upload(ctx *app.UploadFilesContext) error {
	reader, err := ctx.MultipartReader()
	if err != nil {
		return goa.ErrBadRequest("failed to load multipart request: %s", err)
	}
	if reader == nil {
		return goa.ErrBadRequest("not a multipart request")
	}
	//
	var res []string
	convert := func(str string) string {
		hasher := md5.New()
		hasher.Write([]byte(str + time.Now().String()))
		return hex.EncodeToString(hasher.Sum(nil))
	}
	for {
		p, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return goa.ErrBadRequest("failed to load part: %s", err)
		}
		fn := p.FileName()
		filename := convert(fn) + fn[strings.LastIndex(fn, "."):]
		f, err := os.OpenFile("./public/files/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return fmt.Errorf("failed to save file: %s", err) // causes a 500 response
		}
		defer f.Close()
		io.Copy(f, p)
		res = append(res, filename)
	}

	return ctx.OK(res)
}