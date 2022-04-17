package services

import (
	"chitest/pkg/db"
	m "chitest/pkg/models"
	"chitest/pkg/payloads"
	p "chitest/pkg/payloads"
	"net/http"

	"github.com/go-chi/render"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func CreateArticle(w http.ResponseWriter, r *http.Request) {
	data := &p.SSHKeyRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, p.ErrInvalidRequest(err))
		return
	}

	sshKey := data.SSHKey
	sshKey.InsertP(r.Context(), db.DB, boil.Infer())

	render.Status(r, http.StatusCreated)
	render.Render(w, r, p.NewSshKeyResponse(sshKey))
}

func ListSshKeys(w http.ResponseWriter, r *http.Request) {
	keys := m.SSHKeys().AllP(r.Context(), db.DB)
	if err := render.RenderList(w, r, payloads.NewSSHKeyListResponse(keys)); err != nil {
		render.Render(w, r, payloads.ErrRender(err))
		return
	}
}
