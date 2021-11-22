package todo

import (
	"context"
)

type DestroyRequest struct {
	ID string
}

func (t *Todo) Destroy(ctx context.Context, req *DestroyRequest) error {
	return t.kv.Delete(req.ID)
}
