package main

import (
	"log"
	"net/http"
	"time"
)

func CheckAuth() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			log.Println("Checking notification")

			flag := false

			if !flag {
				w.WriteHeader(http.StatusForbidden)

				return
			}

			hf(w, r)
		}
	}
}

func Logging() Middleware {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			defer func() {
				log.Println(r.URL.Path, r.Method, time.Since(start))
			}()

			hf(w, r)
		}
	}
}
