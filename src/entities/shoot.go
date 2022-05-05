package entities

import "fmt"

type Shoot struct {
	TaskNumber        string
	TaskTitle         string
	TrainingFrequency string
}

func (shoot Shoot) ToString() string {
	return fmt.Sprintf("taskNumber: %s\ntaskTitle: %s\nTrainingFrequency: %s", shoot.TaskNumber, shoot.TaskTitle, shoot.TrainingFrequency)
}
