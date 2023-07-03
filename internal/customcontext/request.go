package customcontext

import (
	"context"
)

type key string

var RequestIDKey = key("request_id")

func WithRequestID(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, RequestIDKey, requestID)
}

func RequestID(ctx context.Context) (string, bool) {
	reqID, ok := ctx.Value(RequestIDKey).(string)
	return reqID, ok
}
