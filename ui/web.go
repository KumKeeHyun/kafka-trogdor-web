package ui

import (
	"io"
	"net/http"
	"strings"

	"github.com/KumKeeHyun/kafka-trogdor-web/ui/app"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		f, err := app.Assets.Open("/index.html")
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		defer func() { _ = f.Close() }()
		index, err := io.ReadAll(f)
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}
		c.Writer.Write(index)
	})

	r.GET("/static/*filepath", func(c *gin.Context) {
		c.FileFromFS(strings.TrimPrefix(c.Request.URL.Path, "/static"), app.Assets)
	})
}
