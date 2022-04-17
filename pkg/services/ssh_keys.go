package services

import (
	"chitest/pkg/db"
	"chitest/pkg/models"
	"chitest/pkg/payloads"
	"net/http"

	"github.com/go-chi/render"
)

func ListSshKeys(w http.ResponseWriter, r *http.Request) {
	keys := models.SSHKeys().AllP(r.Context(), db.DB)
	if err := render.RenderList(w, r, payloads.NewSSHKeyListResponse(keys)); err != nil {
		render.Render(w, r, payloads.ErrRender(err))
		return
	}
}
