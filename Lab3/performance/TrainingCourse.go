package performance

const retakeThreshold = 0.7

type TrainingCourse struct {
	Name      string
	Hours     int
	Completed bool
	Score     float64
}

func (c *TrainingCourse) Complete(score float64) {
	c.Completed = true
	c.Score = score
}

func (c TrainingCourse) NeedsRetake() bool {
	return c.Score < retakeThreshold
}

func (c TrainingCourse) DurationHours() int {
	return c.Hours
}
