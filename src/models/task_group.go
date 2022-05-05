package models

import (
	"database/sql"
	"entities"
)

type TaskGroup struct {
	Db *sql.DB
}

func (taskGroup TaskGroup) FindAll() (shoot []entities.Shoot, err error) {
	rows, err := taskGroup.Db.Query("SELECT * FROM task_group.shoot")

	if err != nil {
		return nil, err
	} else {

		var task_group []entities.Shoot
		for rows.Next() {
			var taskNumber string
			var taskTitle string
			var trainingFrequency string

			err2 := rows.Scan(&taskNumber, &taskTitle, &trainingFrequency)
			if err2 != nil {
				return nil, err2
			} else {
				shoot := entities.Shoot{
					TaskNumber:        taskNumber,
					TaskTitle:         taskTitle,
					TrainingFrequency: trainingFrequency,
				}
				task_group = append(task_group, shoot)
			}
		}
		return task_group, nil

	}
}
