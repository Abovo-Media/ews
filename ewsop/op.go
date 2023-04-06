package ewsop

import (
	"context"

	"github.com/Abovo-Media/go-ews"
)

type Operation string

func (op Operation) String() string { return string(op) }

type OperationContext struct {
	Operation Operation
	RequestId string
}

type operationKey struct{}

func setOperation(ctx context.Context, op Operation) context.Context {
	return context.WithValue(ctx, operationKey{}, OperationContext{
		Operation: op,
	})
}

func GetOperation(ctx context.Context) (OperationContext, bool) {
	v, ok := ctx.Value(operationKey{}).(OperationContext)
	if v.RequestId == "" {
		v.RequestId, _ = ews.RequestId(ctx)
	}
	return v, ok
}
