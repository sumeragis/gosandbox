package middleware

import (
	"bytes"
	"io"
	"net/http"

	"github.com/sumeragis/sandbox/backend/logger"
)

func LogHandlingMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		logger.Log.Sugar().Debugf("--> %s %s %s", req.Method, req.URL.Path, req.URL.RawQuery)
		bufOfRequestBody, _ := io.ReadAll(req.Body)
		if body := string(bufOfRequestBody); body != "" {
			logger.Log.Sugar().Debugf("Request Body: %s", string(bufOfRequestBody))
		}
		// 消費されてしまったRequest Bodyを補修する
		req.Body = io.NopCloser(bytes.NewBuffer(bufOfRequestBody))
		lrw := NewLoggingResponseWriter(w)

		next.ServeHTTP(lrw, req)
		statusCode := lrw.statusCode
		logger.Log.Sugar().Debugf("<-- %d %s", statusCode, http.StatusText(statusCode))

	}
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	logger.Log.Sugar().Debugf("Response Body: %v", string(b))
	return lrw.ResponseWriter.Write(b)
}