package repo

import (
	"fmt"
	"github.com/hsmtkk/shiny-happiness/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
)

type sqliteRepo struct {
	db    *gorm.DB
	newID int
}

func NewSqliteRepo(path string) (TaskRepo, error) {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to open database; %s; %w", path, err)
	}
	db.AutoMigrate(&model.Task{})
	return &sqliteRepo{db: db, newID: 1}, nil
}

func (r *sqliteRepo) GetAll() ([]model.Task, error) {
	tasks := []model.Task{}
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

func (r *sqliteRepo) Get(id int) (model.Task, error) {
	var task model.Task
	result := r.db.First(&task, "id = ?", strconv.Itoa(id))
	return task, result.Error
}

func (r *sqliteRepo) New(content string) error {
	result := r.db.Create(&model.Task{ID: r.newID, Content: content})
	if result.Error != nil {
		return result.Error
	}
	r.newID += 1
	return nil
}

func (r *sqliteRepo) Delete(id int) error {
	var task model.Task
	result := r.db.Where("id = ?", strconv.Itoa(id)).Delete(&task)
	return result.Error
}
