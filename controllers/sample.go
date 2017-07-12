package controllers

import (
	"github.com/goadesign/goa"
	"github.com/tikasan/goa-code-reading/app"
)

// SampleController implements the Sample resource.
type SampleController struct {
	*goa.Controller
}

// NewSampleController creates a Sample controller.
func NewSampleController(service *goa.Service) *SampleController {
	return &SampleController{Controller: service.NewController("SampleController")}
}

// GET runs the GET action.
func (c *SampleController) GET(ctx *app.GETSampleContext) error {
	// SampleController_GET: start_implement

	// Put your logic here

	// SampleController_GET: end_implement
	res := app.SampleCollection{}
	return ctx.OK(res)
}

// POST runs the POST action.
func (c *SampleController) POST(ctx *app.POSTSampleContext) error {
	// SampleController_POST: start_implement

	// Put your logic here

	// SampleController_POST: end_implement
	res := &app.Sample{}
	return ctx.OK(res)
}
