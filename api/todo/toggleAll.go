package todo

import (
	"context"
)

type ToggleAllRequest struct {
	ID  string
	Val bool
}
type ToggleAllResponse struct{}

func (t *Todo) ToggleAll(ctx context.Context, req *ToggleAllRequest) (*ToggleAllResponse, error) {
	var items []TodoItem
	if _, err := t.kv.FindAll(&items); err != nil {
		return nil, err
	}
	for _, i := range items {
		i.Completed = req.Val
		if err := t.kv.Put(i.ID, i); err != nil {
			return nil, err
		}
	}
	return nil, nil
}
