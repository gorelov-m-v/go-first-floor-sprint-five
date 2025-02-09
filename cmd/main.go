package main

import (
	"fmt"
	"github.com/Yandex-Practicum/go-1fl-homework-sprint5/internal/workout"
	"time"
)

type CaloriesCalculator interface {
	TrainingInfo() workout.InfoMessage
	Calories() float64
}

func ReadData(training CaloriesCalculator) string {
	return fmt.Sprint(training.TrainingInfo())
}

func main() {
	swimming := workout.Swimming{
		Training: workout.Training{
			TrainingType: "Плавание",
			Action:       2000,
			LenStep:      workout.SwimmingLenStep,
			Duration:     90 * time.Minute,
			Weight:       85,
		},
		LengthPool: 50,
		CountPool:  5,
	}

	fmt.Println(ReadData(swimming))

	walking := workout.Walking{
		Training: workout.Training{
			TrainingType: "Ходьба",
			Action:       20000,
			LenStep:      workout.LenStep,
			Duration:     3*time.Hour + 45*time.Minute,
			Weight:       85,
		},
		Height: 185,
	}

	fmt.Println(ReadData(walking))

	running := workout.Running{
		Training: workout.Training{
			TrainingType: "Бег",
			Action:       5000,
			LenStep:      workout.LenStep,
			Duration:     30 * time.Minute,
			Weight:       85,
		},
	}

	fmt.Println(ReadData(running))
}
