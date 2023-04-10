package app

import (
	"github.com/erry-az/test-go/internal/handler/rest/todo"
	repoTodo "github.com/erry-az/test-go/internal/repo/todo"
	"github.com/erry-az/test-go/internal/server"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Rest(dsn string) error {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	todoR, err := repoTodo.New(db)
	if err != nil {
		return err
	}

	return server.StartRest(todo.New(todoR))
}
