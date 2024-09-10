package main

import (
	"fmt"
	"net/http"
	"testing"
)


func TestNoSurve(t *testing.T) {
	var myH myHandler

	h := NoSurf(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error([]byte(fmt.Sprintf("Type is not http.Handler, but is %T", v)))
	}

	// h := NoSurf(http.AllowQuerySemicolons.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello from middleware")
	// }))
}


func TestSessionLoad(t *testing.T) {
	var myH myHandler

	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		// do nothing
	default:
		t.Error([]byte(fmt.Sprintf("Type is not http.Handler, but is %T", v)))
	}

	// h := NoSurf(http.AllowQuerySemicolons.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintln(w, "Hello from middleware")
	// }))
}
