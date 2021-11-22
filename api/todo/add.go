package todo

import (
	"context"

	"github.com/google/uuid"
)

type AddRequest struct {
	Title string
}
type AddResponse struct{}

func (t *Todo) Add(ctx context.Context, req *AddRequest) (*AddResponse, error) {
	id := uuid.NewString()
	return nil, t.kv.Put(id, &TodoItem{
		ID:        id,
		Title:     req.Title,
		Completed: false,
	})
}
