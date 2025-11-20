package benefits

const claimPending = "pending"
const claimApproved = "approved"
const claimRejected = "rejected"

type InsuranceClaim struct {
	ID          string
	Enrollment  BenefitEnrollment
	Amount      float64
	Approved    bool
	Status      string
	PayoutValue float64
}

func NewInsuranceClaim(id string, enrollment BenefitEnrollment, amount float64) InsuranceClaim {
	return InsuranceClaim{ID: id, Enrollment: enrollment, Amount: amount, Status: claimPending}
}

func (c *InsuranceClaim) Approve() {
	c.Approved = true
	c.Status = claimApproved
	c.PayoutValue = c.Amount
}

func (c *InsuranceClaim) Reject() {
	c.Approved = false
	c.Status = claimRejected
	c.PayoutValue = 0
}

func (c InsuranceClaim) Payout() float64 {
	return c.PayoutValue
}
