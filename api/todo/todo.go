package todo

import (
	"context"

	"github.com/mantil-io/mantil.go"
)

type Todo struct {
	kv *mantil.KV
}

type TodoItem struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type DefaultRequest struct{}
type DefaultResponse struct{}

func New() *Todo {
	kv, _ := mantil.NewKV("todos")
	return &Todo{
		kv: kv,
	}
}

func (t *Todo) Default(ctx context.Context, req *DefaultRequest) (*DefaultResponse, error) {
	panic("not implemented")
}
