// Author: yangzq80@gmail.com
// Date: 2022/8/6
//
package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/yusys-cloud/go-workflow/web/admin/assets"
	"net/http"
	"net/url"
	"os"
)

func adminIndex(ctx *gin.Context) {
	file, err := assets.Static.ReadFile("web/dist/index.html")
	if err != nil && os.IsNotExist(err) {
		ctx.String(http.StatusNotFound, "not found")
		return
	}
	ctx.Data(http.StatusOK, "text/html", file)
}

func handlerStatic(c *gin.Context) {
	staticServer := http.FileServer(http.FS(assets.Static))

	c.Request.URL = &url.URL{Path: "web/dist" + c.Request.RequestURI}

	staticServer.ServeHTTP(c.Writer, c.Request)
}
