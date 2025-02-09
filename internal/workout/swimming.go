package workout

type Swimming struct {
	Training
	LengthPool int
	CountPool  int
}

func (s Swimming) meanSpeed() float64 {
	if s.Duration.Hours() == 0 {
		return 0
	}
	return float64(s.LengthPool*s.CountPool) / MInKm / s.Duration.Hours()
}

func (s Swimming) Calories() float64 {
	return (s.meanSpeed() + SwimmingCaloriesMeanSpeedShift) *
		SwimmingCaloriesWeightMultiplier * s.Weight * s.Duration.Hours()
}

func (s Swimming) TrainingInfo() InfoMessage {
	return InfoMessage{
		TrainingType: s.TrainingType,
		Duration:     s.Duration,
		Distance:     s.distance(),
		Speed:        s.meanSpeed(),
		Calories:     s.Calories(),
	}
}
