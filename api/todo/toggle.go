package todo

import (
	"context"
)

type ToggleRequest struct {
	ID  string
	Val bool
}

func (t *Todo) Toggle(ctx context.Context, req *ToggleRequest) error {
	i := &TodoItem{}
	if err := t.kv.Get(req.ID, i); err != nil {
		return err
	}
	i.Completed = req.Val
	return t.kv.Put(i.ID, i)
}
