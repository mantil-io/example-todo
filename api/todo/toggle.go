package todo

import (
	"context"
)

type ToggleRequest struct {
	ID  string
	Val bool
}
type ToggleResponse struct{}

func (t *Todo) Toggle(ctx context.Context, req *ToggleRequest) (*ToggleResponse, error) {
	i := &TodoItem{}
	if err := t.kv.Get(req.ID, i); err != nil {
		return nil, err
	}
	i.Completed = req.Val
	return nil, t.kv.Put(i.ID, i)
}
