package todo

import (
	"context"
)

type SaveRequest struct {
	Todo TodoItem
}
type SaveResponse struct{}

func (t *Todo) Save(ctx context.Context, req *SaveRequest) (*SaveResponse, error) {
	return nil, t.kv.Put(req.Todo.ID, req.Todo)
}
