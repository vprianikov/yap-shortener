package handlers

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"

	"github.com/vprianikov/yap-shortener/internal/models"
)

func (env *Env) Shorten(c *gin.Context) {
	if ct := c.ContentType(); ct != `text/plain` && ct != `text/plain; charset=utf-8` {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	data, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, errP := url.ParseRequestURI(string(data))
	if errP != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if u.Scheme == `` || u.Host == `` {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	key, errS := env.Storage.Set(models.ExternalURL(data))
	if errS != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.String(http.StatusCreated, `%s/%s`, env.Config.BaseURL(), key)
}
