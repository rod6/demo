package main

import (
	"demo/gateway/ctrl"

	"github.com/arcplus/go-lib/router"
)

func Router() *router.Router {
	r := router.New(MiddlewareLog)
	r.PanicHandler(panicHandler)

	api := r.Group("/api", MiddlewareAuth)
	api.GET("/user/:id", router.WrapX(ctrl.UserMethodX))
	api.GET("/things/y", router.Wrap(ctrl.ThingsMethodY))
	api.GET("/panic", router.Wrap(ctrl.TestPanic))

	return r
}
