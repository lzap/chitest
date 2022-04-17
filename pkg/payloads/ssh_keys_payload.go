package payloads

import (
	"chitest/pkg/models"
	"net/http"

	"github.com/go-chi/render"
)

type SSHKeyPayload struct {
	*models.SSHKey
}
type SSHKeyRequest SSHKeyPayload
type SSHKeyResponse SSHKeyPayload

func (p *SSHKeyRequest) Bind(r *http.Request) error {
	return nil
}

func (p *SSHKeyRequest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (p *SSHKeyResponse) Bind(r *http.Request) error {
	return nil
}

func (p *SSHKeyResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewSshKeyResponse(sshKey *models.SSHKey) render.Renderer {
	return &SSHKeyResponse{SSHKey: sshKey}
}

func NewSSHKeyListResponse(sshKeys []*models.SSHKey) []render.Renderer {
	list := []render.Renderer{}
	for _, k := range sshKeys {
		list = append(list, &SSHKeyResponse{SSHKey: k})
	}
	return list
}

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
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

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
