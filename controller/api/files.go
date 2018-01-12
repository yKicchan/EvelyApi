package api

import (
	"EvelyApi/app"
	"fmt"
	"github.com/goadesign/goa"
	"io"
	"os"
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
	// FilesController_Upload: start_implement

	// Put your logic here
	reader, err := ctx.MultipartReader()
	if err != nil {
		return goa.ErrBadRequest("failed to load multipart request: %s", err)
	}
	if reader == nil {
		return goa.ErrBadRequest("not a multipart request")
	}
	var files []string
	for {
		p, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return goa.ErrBadRequest("failed to load part: %s", err)
		}
		f, err := os.OpenFile("./public/files/"+p.FileName(), os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return fmt.Errorf("failed to save file: %s", err) // causes a 500 response
		}
		defer f.Close()
		io.Copy(f, p)
		files = append(files, f.Name())
	}

	return ctx.OK(files)
	// FilesController_Upload: end_implement
}
