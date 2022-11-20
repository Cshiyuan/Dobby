package future

import "context"

type Task interface {
	Run(ctx context.Context) (interface{}, error)
}


type TaskFunc func(context.Context) (interface{}, error)

func (t TaskFunc) Run(ctx context.Context) (interface{}, error) {
	return t(ctx)
}