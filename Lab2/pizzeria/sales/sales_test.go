package sales_test

import (
	"testing"

	"lab2/pizzeria/communication"
	"lab2/pizzeria/hr"
	"lab2/pizzeria/logistics"
	"lab2/pizzeria/sales"
	"lab2/pizzeria/storage"
)

func TestOrderAndCustomerFlow(t *testing.T) {
	product := logistics.Product{Name: "Pizza"}
	item := sales.OrderItem{Product: product, Quantity: 2, Price: 5}
	if item.Cost() != 10 {
		t.Fatalf("item cost mismatch")
	}
	if item.Describe() == "" {
		t.Fatalf("item describe empty")
	}

	order := sales.Order{Number: "O1", Status: hr.Status{Name: "new"}}
	if err := order.AddItem(item); err != nil {
		t.Fatalf("add item failed: %v", err)
	}
	if err := order.AddItem(sales.OrderItem{Product: product, Quantity: 0}); err == nil {
		t.Fatalf("expected validation error")
	}

	invoice := sales.Invoice{Number: "INV", Order: order}
	invoice.MarkPaid()
	if !invoice.IsPaid() {
		t.Fatalf("invoice should be paid")
	}
	delivery := sales.Delivery{Number: "DEL", Order: order}
	delivery.Dispatch(logistics.Route{Name: "R", Distance: 5})
	delivery.Complete()

	cust := sales.Customer{Name: "Alice", Email: hr.NewEmailAccount("alice@pizza.io")}
	order.Customer = cust
	cust.PlaceOrder(order)
	if cust.OrderCount() != 1 {
		t.Fatalf("order count mismatch")
	}
}

func TestSalesDepartmentAndStorage(t *testing.T) {
	sale := sales.Order{Number: "O1", Status: hr.Status{Name: "new"}}
	sale.Customer = sales.Customer{Name: "Bob", Email: hr.NewEmailAccount("bob@pizza.io")}
	notifier := communication.NewEmailNotifier()
	store := storage.NewMemoryStorage()

	salesDept := sales.SalesDepartment{Storage: store, Notifier: notifier}
	salesDept.RegisterOrder(sale)
	if err := salesDept.PersistOrder(sale); err != nil {
		t.Fatalf("persist order failed: %v", err)
	}
	if _, err := salesDept.FetchOrder("O1"); err != nil {
		t.Fatalf("fetch order failed: %v", err)
	}
	if err := salesDept.NotifyCustomerOrder(sale); err != nil {
		t.Fatalf("notify failed: %v", err)
	}
	if len(notifier.Sent()) == 0 {
		t.Fatalf("expected notification sent")
	}

	delivery := sales.Delivery{Number: "DEL", Order: sale}
	salesDept.Deliveries = append(salesDept.Deliveries, delivery)
	if salesDept.PendingDeliveries() != 1 {
		t.Fatalf("pending deliveries should include undelivered items")
	}
	delivery.Dispatch(logistics.Route{Name: "R1", Distance: 10})
	delivery.Complete()
	salesDept.Deliveries[0] = delivery
	if salesDept.PendingDeliveries() != 0 {
		t.Fatalf("all deliveries should be done")
	}
}

func TestSuppliersAndProcurement(t *testing.T) {
	product := logistics.Product{Name: "Cheese"}
	supplier := sales.Supplier{Name: "Fresh", Materials: []logistics.Material{{Name: "Flour"}}, Products: []logistics.Product{product}}
	if _, err := supplier.ProvideMaterial("Flour"); err != nil {
		t.Fatalf("expected material")
	}
	if _, err := supplier.ProvideProduct("Sauce"); err == nil {
		t.Fatalf("expected supplier error")
	}

	contract := sales.Contract{Number: "C1", Supplier: supplier}
	contract.AddItem("Flour")
	if !contract.Contains("Flour") {
		t.Fatalf("contract should contain item")
	}
	if contract.Contains("x") {
		t.Fatalf("contract should reject missing item")
	}

	po := sales.PurchaseOrder{Number: "PO2", Supplier: supplier}
	po.AddMaterial(logistics.Material{Name: "Flour"})
	if po.MaterialCount() != 1 {
		t.Fatalf("material count mismatch")
	}

	procurement := sales.ProcurementDepartment{Supplier: supplier}
	if err := procurement.ApproveContract(contract); err != nil {
		t.Fatalf("approve failed: %v", err)
	}
	if procurement.OutstandingOrders() != 0 {
		t.Fatalf("no orders yet")
	}
	if err := (&sales.ProcurementDepartment{}).ApproveContract(sales.Contract{}); err == nil {
		t.Fatalf("expected contract error")
	}
}
