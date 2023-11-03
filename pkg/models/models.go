package models

import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: подходящей записи не найдено")

type Task struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Status  bool
}
