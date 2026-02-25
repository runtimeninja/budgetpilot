package mw

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	bytes  int
}

func (w *statusWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = http.StatusOK
	}
	n, err := w.ResponseWriter.Write(b)
	w.bytes += n
	return n, err
}

func RequestLogging(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := &statusWriter{ResponseWriter: w}

			next.ServeHTTP(ww, r)

			lat := time.Since(start)
			reqID := middleware.GetReqID(r.Context())

			logger.Info("http_request",
				"request_id", reqID,
				"method", r.Method,
				"path", r.URL.Path,
				"status", ww.status,
				"bytes", ww.bytes,
				"latency_ms", lat.Milliseconds(),
				"remote_ip", r.RemoteAddr,
				"user_agent", r.UserAgent(),
			)
		})
	}
}
