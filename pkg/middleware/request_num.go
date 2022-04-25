package middleware

import (
	"chitest/pkg/ctxval"
	"context"
	"net/http"
	"sync/atomic"
)

var reqNum uint64

func RequestNum(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		n := atomic.AddUint64(&reqNum, 1)
		ctx = context.WithValue(ctx, ctxval.RequestIdCtxKey, n)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
	return http.HandlerFunc(fn)
}
