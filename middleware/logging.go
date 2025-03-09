package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Define a custom type for context keys.
type contextKey string

// Export the key so other packages can use it.
const HandlerNameKey contextKey = "handlerName"

// LoggingResponseWriter wraps http.ResponseWriter to capture status code and bytes written.
type LoggingResponseWriter struct {
	http.ResponseWriter
	StatusCode int
	Bytes      int64
}

// NewLoggingResponseWriter initializes a new LoggingResponseWriter with default status code 200.
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK, 0}
}

// WriteHeader captures the status code and calls the underlying WriteHeader.
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.StatusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// Write writes data to the connection and tracks the number of bytes written.
func (lrw *LoggingResponseWriter) Write(b []byte) (int, error) {
	n, err := lrw.ResponseWriter.Write(b)
	lrw.Bytes += int64(n)
	return n, err
}

// LoggingMiddleware logs the HTTP method, path, status code, request size, response size, and duration.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		lrw := NewLoggingResponseWriter(w)
		next.ServeHTTP(lrw, r)
		duration := time.Since(start)
		// Retrieve the handler name from context.
		name, ok := r.Context().Value(HandlerNameKey).(string)
		if !ok {
			name = "unknown"
		}
		log.Printf("%-3d %-4s Handler: %-8s RespSize: %-8s Duration: %-8s %-40s ",
			lrw.StatusCode, r.Method, name, formatBytes(lrw.Bytes), formatDuration(duration), r.URL.Path)

	})
}

func formatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func formatDuration(d time.Duration) string {
	if d < time.Second {
		return fmt.Sprintf("%.2fms", float64(d.Nanoseconds())/1e6)
	}
	return fmt.Sprintf("%.2fs", d.Seconds())
}
