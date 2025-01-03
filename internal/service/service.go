package service

import (
	"TaskManager/internal/models"
	"TaskManager/internal/repository"
	"context"
	"errors"
	"time"
)

//type TaskService interface {
//CreateTask(ctx context.Context, title, description string) (*models.Task, error)
//GetAllTasks(ctx context.Context) ([]models.Task, error)
//UpdateTask(ctx context.Context, id int, title, description, status string) error
//DeleteTask(ctx context.Context, id int) error
//}

type TaskService struct {
	repo *repository.TaskRepository
}

func NewTaskService(repo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, title, description string) (*models.Task, error) {
	if title == "" {
		return nil, errors.New("title is empty")
	}

	task := &models.Task{
		Title:       title,
		Description: description,
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	err := s.repo.CreateTask(ctx, *task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (s *TaskService) GetAllTasks(ctx context.Context) ([]models.Task, error) {
	return s.repo.GetTasks(ctx)
}

func (s *TaskService) UpdateTask(ctx context.Context, id int, title, description, status string) error {
	task := &models.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      status,
		UpdateAt:    time.Now(),
	}

	err := s.repo.UpdateTask(ctx, *task)
	if err != nil {
		return err
	}
	return nil
}

func (s *TaskService) DeleteTask(ctx context.Context, id int) error {
	return s.repo.DeleteTask(ctx, id)
}
