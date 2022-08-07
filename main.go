package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yusys-cloud/go-jsonstore-rest/rest"
	"github.com/yusys-cloud/go-workflow/web/admin"
)

func main() {
	r := gin.Default()

	r.Use(admin.Cors(), admin.Exception())

	//Add JSON-REST-API
	rest.NewJsonStoreRest("./data", r)

	admin.ConfigRoutes(r)

	r.Run(":9999")
}
