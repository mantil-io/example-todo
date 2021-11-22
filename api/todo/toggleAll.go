package todo

import (
	"context"
)

type ToggleAllRequest struct {
	ID  string
	Val bool
}

func (t *Todo) ToggleAll(ctx context.Context, req *ToggleAllRequest) error {
	var items []TodoItem
	if _, err := t.kv.FindAll(&items); err != nil {
		return err
	}
	for _, i := range items {
		i.Completed = req.Val
		if err := t.kv.Put(i.ID, i); err != nil {
			return err
		}
	}
	return nil
}
