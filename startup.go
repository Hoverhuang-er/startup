package startup

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"rsc.io/sampler"
)

// Hello returns a friendly greeting.
func Hello() string {
	return sampler.Hello()
}

// RegistryHTTP Registry to HTTP Server with go-chi
func RegistryHTTP() {
	rh := chi.NewMux()
	rh.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
		w.WriteHeader(http.StatusOK)
	})

	if err := http.ListenAndServe(":8080", rh); err != nil {
		log.Fatal(err)
		recover()
	}
}
