package todo

import (
	"context"
)

type DestroyRequest struct {
	ID string
}
type DestroyResponse struct{}

func (t *Todo) Destroy(ctx context.Context, req *DestroyRequest) (*DestroyResponse, error) {
	return nil, t.kv.Delete(req.ID)
}
