package todo

import (
	"context"
)

type GetRequest struct{}
type GetResponse struct {
	Todos []TodoItem
}

func (t *Todo) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	var items []TodoItem
	_, err := t.kv.FindAll(&items)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		Todos: items,
	}, nil
}
