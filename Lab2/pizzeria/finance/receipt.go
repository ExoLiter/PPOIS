package finance

import "fmt"

type Receipt struct {
	Number string
	Items  []string
	Total  float64
}

func (r *Receipt) AddItem(item string, price float64) {
	r.Items = append(r.Items, item)
	r.Total += price
}

func (r Receipt) Summary() string {
	return fmt.Sprintf("receipt:%s:%.2f", r.Number, r.Total)
}
