package middleware

import (
	"chitest/pkg/db"
	m "chitest/pkg/models"
	p "chitest/pkg/payloads"
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func parseInt64(r *http.Request, param string) (int64, error) {
	i, err := strconv.Atoi(chi.URLParam(r, param))
	if err != nil {
		return 0, err
	} else {
		return int64(i), nil
	}
}

func SshKeyCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id int64
		var sshKey *m.SSHKey
		var err error

		if id, err = parseInt64(r, "ID"); err == nil {
			sshKey, err = m.FindSSHKey(r.Context(), db.DB, id)
			if err != nil {
				render.Render(w, r, p.ErrNotFound)
				return
			}
		} else if err != nil {
			render.Render(w, r, p.ErrParamParsingError)
			return
		} else {
			render.Render(w, r, p.ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), "sshKey", sshKey)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
