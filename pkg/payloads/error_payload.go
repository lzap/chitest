package payloads

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

//goland:noinspection ALL
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func (e *ErrResponse) Error() error {
	return errors.New(e.StatusText)
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request",
		ErrorText:      err.Error(),
	}
}

func ErrAWSGeneric(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Error returned from AWS",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found"}
var ErrParamParsingError = &ErrResponse{HTTPStatusCode: 500, StatusText: "Cannot parse parameters"}
var ErrDeleteError = &ErrResponse{HTTPStatusCode: 500, StatusText: "Cannot delete resource"}
