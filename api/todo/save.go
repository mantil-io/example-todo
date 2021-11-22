package todo

import (
	"context"
)

type SaveRequest struct {
	Todo TodoItem
}

func (t *Todo) Save(ctx context.Context, req *SaveRequest) error {
	return t.kv.Put(req.Todo.ID, req.Todo)
}
