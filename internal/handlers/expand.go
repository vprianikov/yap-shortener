package handlers

import (
	"net/http"

	"github.com/vprianikov/yap-shortener/internal/models"
)

func (env *Env) Expand(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	u, err := env.Storage.Get(models.ShortKey(r.PathValue(`shortKey`)))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, string(u), http.StatusTemporaryRedirect)
}
