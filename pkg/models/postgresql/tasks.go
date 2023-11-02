package postgresql

import (
	"database/sql"
	"pilrugen.com/todorugen/pkg/models"
)

// TasksModel - Определяем тип который обертывает пул подключения sql.DB
type TasksModel struct {
	DB *sql.DB
}

func (m *TasksModel) Insert(title, content, created, status string) (int, error) {
	return 0, nil
}

func (m *TasksModel) Get(id int) (*models.Tasks, error) {
	return nil, nil
}
