package marketing

import "lab2/pizzeria/finance"

type MarketingCampaign struct {
	Name    string
	Budget  finance.Budget
	Target  Country
	Channel AdvertisingChannel
	Goals   []string
}

func (m *MarketingCampaign) AddGoal(goal string) {
	m.Goals = append(m.Goals, goal)
}

func (m *MarketingCampaign) Spend(amount float64) error {
	if err := m.Budget.Allocate(m.Name, amount); err != nil {
		return err
	}
	return nil
}
