package performance

const passingScore = 3.0

type PerformanceReview struct {
	ID         string
	EmployeeID string
	Score      float64
	Reviewer   string
	Feedback   string
}

func (r *PerformanceReview) UpdateScore(score float64) {
	r.Score = score
}

func (r PerformanceReview) IsPassing() bool {
	return r.Score >= passingScore
}

func (r *PerformanceReview) AddFeedback(feedback string) {
	r.Feedback = feedback
}
