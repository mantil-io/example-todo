package todo

import (
	"context"

	"github.com/google/uuid"
)

type AddRequest struct {
	Title string
}

func (t *Todo) Add(ctx context.Context, req *AddRequest) error {
	id := uuid.NewString()
	return t.kv.Put(id, &TodoItem{
		ID:        id,
		Title:     req.Title,
		Completed: false,
	})
}
