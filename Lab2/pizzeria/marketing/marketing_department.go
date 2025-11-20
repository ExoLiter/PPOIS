package marketing

import "lab2/pizzeria/finance"

type MarketingDepartment struct {
	Campaigns []MarketingCampaign
	Budget    finance.Budget
}

func (m *MarketingDepartment) Launch(campaign MarketingCampaign) {
	m.Campaigns = append(m.Campaigns, campaign)
}

func (m MarketingDepartment) ActiveCampaigns() int {
	return len(m.Campaigns)
}
