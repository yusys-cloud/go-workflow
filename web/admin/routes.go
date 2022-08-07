// Author: yangzq80@gmail.com
// Date: 2022/8/6
//
package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ConfigRoutes(r *gin.Engine) {

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	ug := r.Group("/api/user")
	{
		ug.POST("/login", login)
		ug.GET("/info", info)
		ug.POST("/logout", logout)
	}

	r.GET("/", adminIndex)
	r.GET("/static/*filepath", handlerStatic)
}
