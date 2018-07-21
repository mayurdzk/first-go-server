package main

import (
	"log"
	"net/http"
)

type LogResponseWriter struct {
	rw http.ResponseWriter
}

func (lrw LogResponseWriter) Header() http.Header {
	return lrw.rw.Header()
}

func (lrw LogResponseWriter) Write(b []byte) (int, error) {
	return lrw.rw.Write(b)
}

func (lrw LogResponseWriter) WriteHeader(statusCode int) {
	log.Printf("Response StatusCode:", statusCode)
	lrw.rw.WriteHeader(statusCode)
}
