package marketing

type CampaignReport struct {
	CampaignName string
	Metrics      map[string]float64
	Notes        []string
}

func (c *CampaignReport) AddMetric(name string, value float64) {
	if c.Metrics == nil {
		c.Metrics = map[string]float64{}
	}
	c.Metrics[name] = value
}

func (c *CampaignReport) AddNote(note string) {
	c.Notes = append(c.Notes, note)
}

func (c CampaignReport) Score() float64 {
	score := 0.0
	for _, v := range c.Metrics {
		score += v
	}
	return score
}
