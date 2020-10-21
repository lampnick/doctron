package mock

import (
	"net/http"
	"net/http/httptest"
)

func HTTPServer(contentType string, expected string, protected bool) *httptest.Server {
	if contentType == "" {
		contentType = "text/html"
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if protected {
			u, p, ok := r.BasicAuth()
			if !ok || (u != "nick" && p != "doctron") {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write([]byte(expected))
	}))
}

func HTTPServerByte(contentType string, expected []byte, protected bool) *httptest.Server {
	if contentType == "" {
		contentType = "text/html"
	}
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if protected {
			u, p, ok := r.BasicAuth()
			if !ok || (u != "nick" && p != "doctron") {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		}

		w.Header().Set("Content-Type", contentType)
		_, _ = w.Write(expected)
	}))
}
