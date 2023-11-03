package main

import "pilrugen.com/todorugen/pkg/models"

type templateData struct {
	Task   *models.Task
	Tasks  []*models.Task
	NextID int
	PrevID int
}
