package services

import (
	"chitest/pkg/clouds/aws"
	"chitest/pkg/db"
	"chitest/pkg/models"
	m "chitest/pkg/models"
	p "chitest/pkg/payloads"
	"net/http"

	"github.com/go-chi/render"
)

func CreateSshKeyResource(w http.ResponseWriter, r *http.Request) {
	existing := r.Context().Value("sshKey").(*m.SSHKey)
	// resource
	cid, err := aws.ImportSSHKey(r.Context(), existing.Body)
	if err != nil {
		render.Render(w, r, p.ErrAWSGeneric(err))
		return
	}
	// db
	res := models.SSHKeyResource{
		Cid: cid,
	}
	existing.AddSSHKeyResourcesP(r.Context(), db.DB, true, &res)
	render.Render(w, r, p.NewSshKeyResponse(existing))
}

func ListSshKeyResources(w http.ResponseWriter, r *http.Request) {
	keys := m.SSHKeyResources().AllP(r.Context(), db.DB)
	if err := render.RenderList(w, r, p.NewSSHKeyResourceListResponse(keys)); err != nil {
		render.Render(w, r, p.ErrRender(err))
		return
	}
}

func DeleteSshKeyResource(w http.ResponseWriter, r *http.Request) {
	existing := r.Context().Value("sshKeyResource").(*m.SSHKeyResource)
	// resource
	err := aws.DeleteSSHKey(r.Context(), existing.Cid)
	if err != nil {
		render.Render(w, r, p.ErrAWSGeneric(err))
		return
	}
	// db
	existing.DeleteP(r.Context(), db.DB)
	resp := p.SSHKeyResourceResponse{
		SSHKeyResource: existing,
	}
	render.Render(w, r, &resp)
}
