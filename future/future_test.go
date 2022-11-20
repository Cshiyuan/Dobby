package future

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func test(ctx context.Context) (interface{}, error) {

	return nil, nil
}

func TestNew(t *testing.T) {

	i := 1
	f := New(context.TODO(), TaskFunc(func(ctx context.Context) (interface{}, error) {
		time.Sleep(time.Second)
		return i, nil
	}))
	r, err := f.Get()
	assert.Equal(t, err, nil)
	assert.Equal(t, r, i)

	f = New(context.TODO(), TaskFunc(func(ctx context.Context) (interface{}, error) {
		panic(errors.New("test"))
		return i, nil
	}))
	r, err = f.Get()
	assert.NotEqual(t, err, nil)


	f = New(context.TODO(), TaskFunc(test))
}
