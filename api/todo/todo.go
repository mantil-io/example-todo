package todo

import (
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

func New() *Todo {
	kv, _ := mantil.NewKV("todos")
	return &Todo{
		kv: kv,
	}
}
