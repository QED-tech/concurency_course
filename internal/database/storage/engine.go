package storage

import (
	"database/internal/database/commands"
	"fmt"
)

type Engine struct {
	storage IStorage
}

func NewEngine(storage IStorage) *Engine {
	return &Engine{
		storage: storage,
	}
}

func (e *Engine) Execute(cmd commands.Command) (Result, error) {
	if err := cmd.Validate(); err != nil {
		return Result{}, err
	}

	switch cmd.Operation {
	case commands.SetOperation:
		return e.storage.Set(
			cmd.GetKey(),
			cmd.GetValue(),
		), nil
	case commands.GetOperation:
		return e.storage.Get(
			cmd.GetKey(),
		), nil
	case commands.DeleteOperation:
		return e.storage.Delete(
			cmd.GetKey(),
		), nil
	}

	return Result{}, fmt.Errorf("unsupported operation %s", cmd.Operation)
}
