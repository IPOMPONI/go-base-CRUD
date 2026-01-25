package utils

import "context"

type bookIdType struct{}

var bookIdKey = bookIdType{}

func SetBookIdInCtx(ctx context.Context, id int) context.Context {
	return context.WithValue(ctx, bookIdKey, id)
}

func GetBookIdFromCtx(ctx context.Context) (int, bool) {
	id, ok := ctx.Value(bookIdKey).(int)
	return id, ok
}
