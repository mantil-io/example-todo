package todo

import (
	"context"
)

func (t *Todo) ClearCompleted(ctx context.Context) error {
	var items []TodoItem
	if _, err := t.kv.FindAll(&items); err != nil {
		return err
	}
	var toDelete []string
	for _, i := range items {
		if i.Completed {
			toDelete = append(toDelete, i.ID)
		}
	}
	if err := t.kv.Delete(toDelete...); err != nil {
		return err
	}
	return nil
}
