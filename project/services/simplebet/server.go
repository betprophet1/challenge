package simplebet

import (
	"fmt"
	"net/http"

	"project/app"
	"project/app/lib/gin"
	"project/app/svcprovider"
)

func StartSimpleBet(host string, port int) {
	router := gin.New()

	InitRouter(router)
	session := app.NewSession(svcprovider.NewHttpService(
		&http.Server{
			Handler: router,
			Addr:    fmt.Sprintf("%s:%d", host, port),
		},
	))
	app.RunSession(session)
}
