package handle

import (
	"github.com/hsmtkk/shiny-happiness/model"
	"github.com/hsmtkk/shiny-happiness/repo"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
)

type Handler interface {
	GetAll(ctx echo.Context) error
	Get(ctx echo.Context) error
	EditGet(ctx echo.Context) error
	EditPost(ctx echo.Context) error
	DeleteGet(ctx echo.Context) error
	DeletePost(ctx echo.Context) error
	NewGet(ctx echo.Context) error
	NewPost(ctx echo.Context) error
}

type handlerImpl struct {
	taskRepo repo.TaskRepo
}

func New(taskRepo repo.TaskRepo) Handler {
	return &handlerImpl{taskRepo: taskRepo}
}

func (h *handlerImpl) GetAll(ctx echo.Context) error {
	tasks, err := h.taskRepo.GetAll()
	if err != nil {
		log.Print(err.Error())
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.Render(http.StatusOK, "getall", tasks)
}

func (h *handlerImpl) Get(ctx echo.Context) error {
	task, err := h.getTaskFromID(ctx)
	if err != nil {
		log.Print(err.Error())
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.Render(http.StatusOK, "get", task)
}

func (h *handlerImpl) EditGet(ctx echo.Context) error {
	return nil
}

func (h *handlerImpl) EditPost(ctx echo.Context) error {
	return nil
}

func (h *handlerImpl) DeleteGet(ctx echo.Context) error {
	task, err := h.getTaskFromID(ctx)
	if err != nil {
		log.Print(err.Error())
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.Render(http.StatusOK, "delete", task)
}

func (h *handlerImpl) DeletePost(ctx echo.Context) error {
	taskID, err := h.getTaskID(ctx)
	if err != nil {
		log.Print(err.Error())
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	if err := h.taskRepo.Delete(taskID); err != nil {
		log.Print(err.Error())
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.Redirect(http.StatusMovedPermanently, "/todo")
}

func (h *handlerImpl) NewGet(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "new", nil)
}

func (h *handlerImpl) NewPost(ctx echo.Context) error {
	content := ctx.FormValue("content")
	if err := h.taskRepo.New(content); err != nil {
		log.Print(err.Error())
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.Redirect(http.StatusMovedPermanently, "/todo")
}

func (h *handlerImpl) getTaskID(ctx echo.Context) (int, error) {
	taskID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return 0, err
	}
	return taskID, nil
}

func (h *handlerImpl) getTaskFromID(ctx echo.Context) (model.Task, error) {
	taskID, err := h.getTaskID(ctx)
	if err != nil {
		return model.Task{}, err
	}
	task, err := h.taskRepo.Get(taskID)
	if err != nil {
		return model.Task{}, err
	}
	return task, nil
}
