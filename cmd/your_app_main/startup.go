// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package startup provides a simple HTTP server that can be used to
package startup // import "github.com/Hoverhuang-er/startup"

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
	// Register routes use chi
	rh := chi.NewMux()
	// Register routes / to say welcome
	rh.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
		w.WriteHeader(http.StatusOK)
	})
	// Register routes /hello to say hello at :8080
	if err := http.ListenAndServe(":8080", rh); err != nil {
		log.Fatal(err)
		recover()
	}
}
