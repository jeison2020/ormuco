package middlewares

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type ResponseRecorder struct {
	http.ResponseWriter
	StatusCode int
}

func (respr *ResponseRecorder) WriteHeader(statusCode int) {
	respr.StatusCode = statusCode
	respr.ResponseWriter.WriteHeader(statusCode)
}


func Logger(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		responseRecorder := &ResponseRecorder{
			ResponseWriter: res,
		}
		handler.ServeHTTP(responseRecorder, req)
		duration := time.Since(startTime)
		logger := log.Info()
		if responseRecorder.StatusCode != http.StatusOK {
			logger = log.Error()
		}
		logger.Str("protocol", "http").
			Str("method", req.Method).
			Str("path", req.RequestURI).
			Int("status_code", responseRecorder.StatusCode).
			Str("status_text", http.StatusText(responseRecorder.StatusCode)).
			Dur("duration", duration).
			Msg("received a HTTP request")
	})

}
