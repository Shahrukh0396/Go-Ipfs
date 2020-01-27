package handlers

import (
	"log"
	"net/http"
	"time"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Something went wrong. \n"))
		return
	} else {
		t1 := time.Now()
		t2 := time.Now()
		log.Printf(r.URL.Path)
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t2.Sub(t1))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Welcome to IPFS Server. \n"))
	}

}
