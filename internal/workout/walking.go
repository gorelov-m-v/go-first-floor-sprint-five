package workout

type Walking struct {
	Training
	Height float64
}

func (w Walking) Calories() float64 {
	return (CaloriesWeightMultiplier*w.Weight +
		((w.meanSpeed()*KmHInMsec*w.meanSpeed()*KmHInMsec)/
			(w.Height/CmInM))*CaloriesSpeedHeightMultiplier*w.Weight) *
		w.Duration.Hours() * MinInHours
}

func (w Walking) TrainingInfo() InfoMessage {
	info := w.Training.TrainingInfo()
	info.Calories = w.Calories()
	return info
}
