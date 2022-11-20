// Package future TODO
package future

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"

)

// Future Future最简单实现
type Future interface {
	Get() (interface{}, error)
}


type futureImpl struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

// New 创建一个Future
func New(ctx context.Context, t Task) Future {
	f := futureImpl{
		wg: sync.WaitGroup{},
	}
	f.wg.Add(1)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				stack := getStack()
				f.err = errors.New(fmt.Sprintf("safe go routine panic, err: %v, stack: %v", err, stack))
			}
			f.wg.Done()
		}()
		f.val, f.err = t.Run(ctx)
	}()
	return &f
}

// Get 获得数据
func (f *futureImpl) Get() (interface{}, error) {
	f.wg.Wait()
	return f.val, f.err
}

// getStack get runtime stack
func getStack() string {
	buf := make([]byte, 1024*16)
	n := runtime.Stack(buf, true)
	stack := buf[:n]
	str := fmt.Sprintf("%s", stack)
	return str
}
