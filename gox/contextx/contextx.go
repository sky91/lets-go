package contextx

import "context"

func WithValueFrom(parent context.Context, valueFrom context.Context) context.Context {
	return valueFromContext{Context: parent, valueFrom: valueFrom}
}

type valueFromContext struct {
	context.Context
	valueFrom context.Context
}

func (thisV valueFromContext) Value(key any) any {
	return thisV.valueFrom.Value(key)
}
