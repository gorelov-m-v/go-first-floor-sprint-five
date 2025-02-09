package workout

type Running struct {
	Training
}

func (r Running) Calories() float64 {
	return (CaloriesMeanSpeedMultiplier*r.meanSpeed() + CaloriesMeanSpeedShift) *
		r.Weight / MInKm * r.Duration.Hours() * MinInHours
}

func (r Running) TrainingInfo() InfoMessage {
	info := r.Training.TrainingInfo()
	info.Calories = r.Calories()
	return info
}
