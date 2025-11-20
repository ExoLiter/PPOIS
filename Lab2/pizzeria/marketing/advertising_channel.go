package marketing

type AdvertisingChannel struct {
	Name string
	Cost float64
	Ads  []Advertisement
}

func (c *AdvertisingChannel) Schedule(ad Advertisement) {
	c.Ads = append(c.Ads, ad)
}

func (c AdvertisingChannel) TotalAds() int {
	return len(c.Ads)
}
