package production_test

import (
	"testing"

	"lab2/pizzeria/logistics"
	"lab2/pizzeria/production"
)

func TestProductionFlow(t *testing.T) {
	product := logistics.Product{Name: "Pizza"}
	material := logistics.Material{Name: "Flour"}
	product.AddIngredient(material)

	card := production.TechnologicalCard{Name: "PizzaCard", Product: product}
	card.AddStep("Mix")
	if card.StepCount() == 0 {
		t.Fatalf("step count missing")
	}

	order := production.ProductionOrder{Number: "PO1", Card: card, Product: product, Quantity: 5}
	order.Schedule(3)
	if order.Remaining() != 2 {
		t.Fatalf("remaining mismatch")
	}

	unit := production.ProductionUnit{Name: "Oven", Product: product, Capacity: 10}
	unit.AssignOrder(order)
	unit.Toggle(true)

	line := production.ProductionLine{Name: "Line1", Unit: unit}
	line.Enqueue(order)

	plan := production.ProductionPlan{Name: "Plan", Materials: []logistics.Material{{Name: "Cheese"}}}
	plan.AddOrder(order)
	if !plan.NeedsMaterial("Cheese") {
		t.Fatalf("material missing")
	}

	factory := production.Factory{Name: "Central", Lines: []production.ProductionLine{line}, Plan: plan}
	if factory.Produce() <= 0 {
		t.Fatalf("produce should be positive")
	}
}

func TestProductionMetrics(t *testing.T) {
	line := production.ProductionLine{}
	line.Enqueue(production.ProductionOrder{Quantity: 3})
	if line.Load() != 3 {
		t.Fatalf("line load mismatch")
	}

	factory := production.Factory{}
	factory.AddLine(line)
	if factory.Produce() != 3 {
		t.Fatalf("factory produce mismatch")
	}
}
