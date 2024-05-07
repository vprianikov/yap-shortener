package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/vprianikov/yap-shortener/internal/models"
)

func (env *Env) Expand(c *gin.Context) {
	u, err := env.Storage.Get(models.ShortKey(c.Param(`shortKey`)))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, string(u))
}
