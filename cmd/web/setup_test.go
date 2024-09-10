package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M)  {
	os.Exit(m.Run())	
}

// crate struct for myHandler
type myHandler struct {

}

func (mh *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	

	
}