package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func router() http.Handler {
	mux := http.NewServeMux()

	purl, _ := url.Parse("http://localhost:5173")
	rproxy := httputil.NewSingleHostReverseProxy(purl)

	mux.HandleFunc("/", rproxy.ServeHTTP)

	return mux
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	srv := &http.Server{
		Addr:    ":8888",
		Handler: logRequest(router()),
	}

	fmt.Println("Server listening on port 8888")

	log.Panic(
		srv.ListenAndServe(),
	)
}
