package repo

import (
	"github.com/hsmtkk/shiny-happiness/model"
)

type TaskRepo interface {
	GetAll() ([]model.Task, error)
	Get(id int) (model.Task, error)
	New(content string) error
	Delete(id int) error
}
