package todo

import (
	"context"
	"github.com/erry-az/test-go/internal/entity"
	"gorm.io/gorm"
)

type Repo struct {
	db *gorm.DB
}

func New(db *gorm.DB) (*Repo, error) {
	err := db.AutoMigrate(&entity.Todo{})
	if err != nil {
		return nil, err
	}

	return &Repo{db: db}, nil
}

func (r *Repo) Get(ctx context.Context, id uint) (entity.Todo, error) {
	var todo entity.Todo

	get := r.db.WithContext(ctx).First(&todo, id)

	return todo, get.Error
}

func (r *Repo) Create(ctx context.Context, todo entity.Todo) (uint, error) {
	create := r.db.WithContext(ctx).Create(&todo)

	return todo.ID, create.Error
}
