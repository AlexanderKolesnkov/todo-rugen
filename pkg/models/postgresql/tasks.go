package postgresql

import (
	"database/sql"
	"errors"
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
	stmt := `SELECT id, title, content, created, status FROM tasks
	WHERE id = $1`

	t := &models.Tasks{}

	err := m.DB.QueryRow(stmt, id).Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return t, nil
}
