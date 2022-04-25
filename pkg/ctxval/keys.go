package ctxval

import "context"

type CommonKeyId int

const (
	LoggerCtxKey         CommonKeyId = iota
	RequestIdCtxKey      CommonKeyId = iota
	SshKeyCtxKey         CommonKeyId = iota
	SshKeyResourceCtxKey CommonKeyId = iota
)

func GetValue[T string](ctx context.Context, key CommonKeyId) T {
	return ctx.Value(key).(T)
}
