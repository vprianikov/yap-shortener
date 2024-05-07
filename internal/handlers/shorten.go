package handlers

import (
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/vprianikov/yap-shortener/internal/models"
)

func readRequest(r *http.Request) (string, error) {
	if r.Header.Get(`Content-Type`) != `text/plain` && r.Header.Get(`Content-Type`) != `text/plain; charset=utf-8` {
		return ``, errors.New(`unsupported media type`)
	}

	body, errR := io.ReadAll(r.Body)
	if errR != nil {
		return ``, errR
	}

	u, errP := url.ParseRequestURI(string(body))
	if errP != nil {
		return ``, errP
	}
	if u.Scheme == `` || u.Host == `` {
		return ``, errors.New(`scheme and host must be defined`)
	}

	return string(body), nil
}

func (env *Env) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	u, errR := readRequest(r)
	if errR != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	key, errS := env.Storage.Set(models.ExternalURL(u))
	if errS != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set(`Content-Type`, `text/plain; charset=utf-8`)
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte(`http://` + env.Config.Host() + `:` + env.Config.Port() + `/` + string(key)))
}
