package middlewares

import (
	"net/http"

	"github.com/bmviniciuss/cccat12/internal/customcontext"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

var RequestIDHeader = "X-Request-Id"

func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID := r.Header.Get(RequestIDHeader)
		if reqID == "" {
			reqID = entities.NewULID().String()
		}
		ctx = customcontext.WithRequestID(ctx, reqID)
		w.Header().Set(RequestIDHeader, reqID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
