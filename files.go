package main

import (
	"EvelyApi/app"
	"github.com/goadesign/goa"
)

// FilesController implements the files resource.
type FilesController struct {
	*goa.Controller
}

// NewFilesController creates a files controller.
func NewFilesController(service *goa.Service) *FilesController {
	return &FilesController{Controller: service.NewController("FilesController")}
}

// Upload runs the upload action.
func (c *FilesController) Upload(ctx *app.UploadFilesContext) error {
	// FilesController_Upload: start_implement

	// Put your logic here

	return nil
	// FilesController_Upload: end_implement
}
