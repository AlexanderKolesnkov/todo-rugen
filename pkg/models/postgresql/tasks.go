package postgresql

import (
	"database/sql"
	"errors"
	"pilrugen.com/todorugen/pkg/models"
)

// TaskModel - Определяем тип который обертывает пул подключения sql.DB
type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) Insert(title, content, created, status string) (int, error) {
	return 0, nil
}

func (m *TaskModel) Get(id int) (*models.Task, error) {
	stmt := `SELECT id, title, content, created, status FROM tasks
	WHERE id = $1`

	t := &models.Task{}

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

func (m *TaskModel) GetAll() ([]*models.Task, error) {
	stmt := `SELECT * FROM tasks`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []*models.Task

	for rows.Next() {
		t := &models.Task{}

		err = rows.Scan(&t.ID, &t.Title, &t.Content, &t.Created, &t.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (m *TaskModel) MaxID() (int, error) {
	stmt := `SELECT MAX(id) FROM tasks`

	var id int
	err := m.DB.QueryRow(stmt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
