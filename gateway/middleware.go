package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/arcplus/go-lib/log"
)

func panicHandler(rw http.ResponseWriter, r *http.Request, x interface{}) {
	logger := log.Logger()
	if tid := r.Header.Get("x-request-id"); tid != "" {
		logger = logger.Trace(tid)
	}

	raw, _ := httputil.DumpRequest(r, false)

	logger.WithStack().Errorf("recover http: %s\nerr: %v\n", string(raw), x)

	rw.WriteHeader(http.StatusInternalServerError)
	rw.Write([]byte(fmt.Sprint(x)))
}

func MiddlewareLog(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	buf := &bytes.Buffer{}

	ctx := context.WithValue(r.Context(), "x-ctx-body", buf)
	r = r.WithContext(ctx)
	next(rw, r)
}

func MiddlewareAuth(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)
}
