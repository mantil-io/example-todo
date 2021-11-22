package todo

import (
	"context"
)

type GetResponse struct {
	Todos []TodoItem
}

func (t *Todo) Get(ctx context.Context) (*GetResponse, error) {
	var items []TodoItem
	_, err := t.kv.FindAll(&items)
	if err != nil {
		return nil, err
	}
	return &GetResponse{
		Todos: items,
	}, nil
}
