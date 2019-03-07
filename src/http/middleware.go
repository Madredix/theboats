package http

import (
	"context"
	chiMiddleware "github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

const (
	RequestID = iota
	LoggerID
)

func Logger(logger *logrus.Entry) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			timeStart := time.Now()
			ww := chiMiddleware.NewWrapResponseWriter(w, r.ProtoMajor)

			requestId := r.Header.Get("Request-ID")
			if reqID := r.Context().Value(chiMiddleware.RequestIDKey); requestId == "" && reqID != nil {
				requestId = reqID.(string)
				requestId = string(requestId[len(requestId)-17:])
			}
			ctxLogger := logger.WithField(`request-id`, requestId)
			ctx := context.WithValue(r.Context(), LoggerID, ctxLogger)
			ctx = context.WithValue(ctx, RequestID, requestId)
			w.Header().Set("Request-ID", requestId)
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(ww, r.WithContext(ctx))
			ctxLogger.WithFields(logrus.Fields{
				`action`:   `request completed`,
				`status`:   ww.Status(),
				`method`:   r.Method,
				`url`:      r.RequestURI,
				`duration`: int(time.Since(timeStart) / time.Millisecond)},
			).Debug()
		})
	}
}

func Recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {

				var errStr string
				switch x := rvr.(type) {
				case string:
					errStr = x
				case error:
					errStr = x.Error()
				default:
					errStr = "Unknown panic"
				}

				logger := r.Context().Value(LoggerID).(*logrus.Entry)
				logger.WithField(`action`, `request completed`).Error(errStr)

				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, `{"status": 500, "data": "INTERNAL_ERROR"}`) // nolint:errcheck
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
