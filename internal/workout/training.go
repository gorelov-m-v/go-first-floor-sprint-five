package workout

import (
	"fmt"
	"time"
)

type InfoMessage struct {
	TrainingType string
	Duration     time.Duration
	Distance     float64
	Speed        float64
	Calories     float64
}

func (i InfoMessage) String() string {
	return fmt.Sprintf(
		"Тип тренировки: %s\n"+
			"Длительность: %v мин\n"+
			"Дистанция: %.2f км.\n"+
			"Ср. скорость: %.2f км/ч\n"+
			"Потрачено ккал: %.2f\n",
		i.TrainingType,
		i.Duration.Minutes(),
		i.Distance,
		i.Speed,
		i.Calories,
	)
}

type Training struct {
	TrainingType string
	Action       int
	LenStep      float64
	Duration     time.Duration
	Weight       float64
}

func (t Training) distance() float64 {
	return float64(t.Action) * t.LenStep / MInKm
}

func (t Training) meanSpeed() float64 {
	if t.Duration.Hours() == 0 {
		return 0
	}
	return t.distance() / t.Duration.Hours()
}

func (t Training) Calories() float64 {
	return 0
}

func (t Training) TrainingInfo() InfoMessage {
	return InfoMessage{
		TrainingType: t.TrainingType,
		Duration:     t.Duration,
		Distance:     t.distance(),
		Speed:        t.meanSpeed(),
		Calories:     t.Calories(),
	}
}
