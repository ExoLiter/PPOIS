package marketing_test

import (
	"testing"

	"lab2/pizzeria/finance"
	"lab2/pizzeria/marketing"
)

func TestMarketingFlow(t *testing.T) {
	country := marketing.Country{Name: "Italy", TaxRate: 0.2}
	country.AddRegulation("pizza-size")
	if country.EffectiveTax(100) != 20 {
		t.Fatalf("tax mismatch")
	}

	ad := marketing.Advertisement{Title: "Margherita", Message: "Fresh"}
	ad.AddChannel("social")
	if ad.Content() == "" {
		t.Fatalf("ad content empty")
	}
	channel := marketing.AdvertisingChannel{Name: "Online"}
	channel.Schedule(ad)
	if channel.TotalAds() != 1 {
		t.Fatalf("ads not tracked")
	}

	campaign := marketing.MarketingCampaign{Name: "Launch", Budget: finance.NewBudget("Mkt", 50), Target: country, Channel: channel}
	campaign.AddGoal("awareness")
	if err := campaign.Spend(60); err == nil {
		t.Fatalf("expected overspend")
	}
}

func TestMarketingReporting(t *testing.T) {
	campaignReport := marketing.CampaignReport{CampaignName: "Launch"}
	campaignReport.AddMetric("reach", 10)
	campaignReport.AddNote("good")
	if campaignReport.Score() <= 0 {
		t.Fatalf("report score incorrect")
	}

	campaign := marketing.MarketingCampaign{Name: "Extra", Budget: finance.NewBudget("Dept", 20), Target: marketing.Country{Name: "USA"}}
	campaign.AddGoal("goal")
	mDept := marketing.MarketingDepartment{Budget: finance.NewBudget("MD", 40)}
	mDept.Launch(campaign)
	if mDept.ActiveCampaigns() != 1 {
		t.Fatalf("campaign count mismatch")
	}
}
