package todo

import (
	"context"
)

type ClearCompletedRequest struct{}
type ClearCompletedResponse struct{}

func (t *Todo) ClearCompleted(ctx context.Context, req *ClearCompletedRequest) (*ClearCompletedResponse, error) {
	var items []TodoItem
	if _, err := t.kv.FindAll(&items); err != nil {
		return nil, err
	}
	for _, i := range items {
		if i.Completed {
			if err := t.kv.Delete(i.ID); err != nil {
				return nil, err
			}
		}
	}
	return nil, nil
}
