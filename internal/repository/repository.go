package repository

import (
	"TaskManager/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TaskRepository struct {
	db *pgxpool.Pool
}

func NewTaskRepository(db *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task models.Task) error {

	_, err := r.db.Exec(ctx,
		"INSERT INTO tasks (title, description, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)",
		task.Title, task.Description, task.Status, task.CreatedAt, task.UpdateAt)
	return err
}

func (r *TaskRepository) GetTasks(ctx context.Context) ([]models.Task, error) {
	var tasks []models.Task
	rows, err := r.db.Query(ctx, "SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var t models.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.CreatedAt, &t.UpdateAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}

func (r *TaskRepository) UpdateTask(ctx context.Context, task models.Task) error {
	_, err := r.db.Exec(ctx,
		"UPDATE tasks SET title=$1, description=$2, status=$3, updated_at=$4 WHERE id=$5",
		task.Title, task.Description, task.Status, task.UpdateAt, task.ID)

	return err
}

func (r *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	_, err := r.db.Exec(ctx,
		"DELETE FROM tasks WHERE id=$1", id)
	return err
}
