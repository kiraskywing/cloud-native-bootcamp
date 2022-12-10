package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func getClientIP(r *http.Request) string {
	resultIP := r.RemoteAddr
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		resultIP = ip
	} else if ip = r.Header.Get("X-Forwarded-For"); ip != "" {
		resultIP = ip
	} else {
		resultIP, _, _ = net.SplitHostPort(resultIP)
	}

	if resultIP == "::1" {
		resultIP = "127.0.0.1"
	}

	return resultIP
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	clientIP := getClientIP(r)
	_, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("client IP: %s, status code: %d\n", clientIP, http.StatusBadRequest)
	}

	for k, v := range r.Header {
		if k == "Content-Length" {
			continue
		}

		for _, sub_v := range v {
			w.Header().Add(k, sub_v)
		}
	}

	version := os.Getenv("VERSION")
	if len(version) > 0 {
		w.Header().Add("VERSION", version)
	}

	log.Printf("== client IP: %s, status code: %d ==\n", clientIP, http.StatusOK)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
