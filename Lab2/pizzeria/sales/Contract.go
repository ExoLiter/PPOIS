package sales

type Contract struct {
	Number   string
	Supplier Supplier
	Items    []string
}

func (c *Contract) AddItem(item string) {
	c.Items = append(c.Items, item)
}

func (c Contract) Contains(item string) bool {
	for _, entry := range c.Items {
		if entry == item {
			return true
		}
	}
	return false
}
