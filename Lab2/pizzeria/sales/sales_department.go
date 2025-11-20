package sales

import (
	"fmt"

	"lab2/pizzeria/communication"
	"lab2/pizzeria/storage"
)

type SalesDepartment struct {
	Orders     []Order
	Customers  []Customer
	Invoices   []Invoice
	Deliveries []Delivery
	Items      []OrderItem
	Storage    storage.Storage
	Notifier   communication.Notifier
}

func (s *SalesDepartment) RegisterOrder(order Order) {
	s.Orders = append(s.Orders, order)
}

func (s SalesDepartment) PendingDeliveries() int {
	pending := 0
	for _, delivery := range s.Deliveries {
		if !delivery.Delivered {
			pending++
		}
	}
	return pending
}

func (s *SalesDepartment) PersistOrder(order Order) error {
	if s.Storage == nil {
		return fmt.Errorf("storage not configured")
	}
	return s.Storage.Save(order.Number, order)
}

func (s *SalesDepartment) FetchOrder(number string) (Order, error) {
	if s.Storage == nil {
		return Order{}, fmt.Errorf("storage not configured")
	}
	v, err := s.Storage.Load(number)
	if err != nil {
		return Order{}, err
	}
	order, ok := v.(Order)
	if !ok {
		return Order{}, fmt.Errorf("invalid order type")
	}
	return order, nil
}

func (s *SalesDepartment) NotifyCustomerOrder(order Order) error {
	if s.Notifier == nil || order.Customer.Name == "" {
		return fmt.Errorf("notifier not configured")
	}
	return s.Notifier.Notify(order.Customer.Name, fmt.Sprintf("order %s status %s", order.Number, order.Status.Name))
}
