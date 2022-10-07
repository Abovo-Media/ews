package ewsop

import (
	"context"
)

type Operation string

func (op Operation) String() string { return string(op) }

type operationKey struct{}

func setOperation(ctx context.Context, op Operation) context.Context {
	return context.WithValue(ctx, operationKey{}, op)
}

func GetOperation(ctx context.Context) (Operation, bool) {
	v, ok := ctx.Value(operationKey{}).(Operation)
	return v, ok
}
