package repo

import (
	"fmt"
	"github.com/hsmtkk/shiny-happiness/model"
)

type memoryRepo struct {
	tasks map[int]model.Task
	newID int
}

func NewMemoryRepo() TaskRepo {
	tasks := map[int]model.Task{}
	tasks[1] = model.Task{ID: 1, Content: "alpha"}
	tasks[2] = model.Task{ID: 2, Content: "bravo"}
	return &memoryRepo{tasks: tasks, newID: 3}
}

func (m *memoryRepo) GetAll() ([]model.Task, error) {
	tasks := []model.Task{}
	for _, task := range m.tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (m *memoryRepo) Get(id int) (model.Task, error) {
	task, ok := m.tasks[id]
	if ok {
		return task, nil
	}
	return model.Task{}, fmt.Errorf("user ID %d is not found", id)
}

func (m *memoryRepo) New(content string) error {
	m.tasks[m.newID] = model.Task{ID: m.newID, Content: content}
	m.newID += 1
	return nil
}

func (m *memoryRepo) Delete(taskID int) error {
	delete(m.tasks, taskID)
	return nil
}
