package todo

import (
	"context"
)

func (t *Todo) ClearCompleted(ctx context.Context) error {
	var items []TodoItem
	if _, err := t.kv.FindAll(&items); err != nil {
		return err
	}
	for _, i := range items {
		if i.Completed {
			if err := t.kv.Delete(i.ID); err != nil {
				return err
			}
		}
	}
	return nil
}
