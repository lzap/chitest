package services

import (
	"chitest/pkg/db"
	m "chitest/pkg/models"
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
	if err := render.RenderList(w, r, p.NewSSHKeyListResponse(keys)); err != nil {
		render.Render(w, r, p.ErrRender(err))
		return
	}
}

func GetSshKey(w http.ResponseWriter, r *http.Request) {
	sshKey := r.Context().Value("sshKey").(*m.SSHKey)
	render.Render(w, r, p.NewSshKeyResponse(sshKey))
}

func DeleteSshKey(w http.ResponseWriter, r *http.Request) {
	sshKey := r.Context().Value("sshKey").(*m.SSHKey)
	rows := sshKey.DeleteP(r.Context(), db.DB)
	if rows == 1 {
		render.Render(w, r, p.NewSshKeyResponse(sshKey))
	} else {
		render.Render(w, r, p.ErrDeleteError)
	}
}

func UpdateSshKey(w http.ResponseWriter, r *http.Request) {
	existing := r.Context().Value("sshKey").(*m.SSHKey)
	data := &p.SSHKeyRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, p.ErrInvalidRequest(err))
		return
	}

	updated := data.SSHKey
	existing.Body = updated.Body

	rows := existing.UpdateP(r.Context(), db.DB, boil.Infer())
	if rows == 1 {
		render.Render(w, r, p.NewSshKeyResponse(existing))
	} else {
		render.Render(w, r, p.ErrDeleteError)
	}
}
