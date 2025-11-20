package logistics_test

import (
	"testing"

	"lab2/pizzeria/logistics"
	"lab2/pizzeria/marketing"
)

func TestLogisticsFlow(t *testing.T) {
	country := marketing.Country{Name: "Italy", TaxRate: 0.2}
	warehouse := logistics.Warehouse{Name: "Main", Country: country, Capacity: 10}
	material := logistics.Material{Name: "Cheese", Unit: "kg", Stock: 5}
	if err := material.Consume(10); err == nil {
		t.Fatalf("expected shortage")
	}
	material.Restock(5)
	cargo := logistics.Cargo{ID: "C1", Weight: 3, Material: &material}
	if err := warehouse.Store(cargo); err != nil {
		t.Fatalf("store failed: %v", err)
	}
	if warehouse.Utilization() <= 0 {
		t.Fatalf("expected utilization")
	}

	route := logistics.Route{Name: "R1", Distance: 10}
	if _, err := route.TravelTime(0); err == nil {
		t.Fatalf("expected route error")
	}

	vehicle := logistics.Vehicle{Plate: "V1", Type: logistics.VehicleType{Name: "Van", Capacity: 5}}
	if err := vehicle.Load(cargo); err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if err := vehicle.Load(logistics.Cargo{ID: "Heavy", Weight: 10}); err == nil {
		t.Fatalf("expected overload")
	}
	if vehicle.LoadCount() == 0 {
		t.Fatalf("load count missing")
	}

	shipment := logistics.Shipment{ID: "S1", Route: route, Vehicle: vehicle}
	shipment.AddCargo(cargo)
	if shipment.Summary() == "" {
		t.Fatalf("shipment summary empty")
	}

	dept := logistics.LogisticsDepartment{Country: country}
	if err := dept.Dispatch(shipment); err != nil {
		t.Fatalf("dispatch failed: %v", err)
	}
	if dept.TotalShipments() != 1 {
		t.Fatalf("total shipments mismatch")
	}
}

func TestCargoPlansAndRouting(t *testing.T) {
	plan := logistics.CargoSortingPlan{Name: "Plan"}
	plan.AddCargo(logistics.Cargo{ID: "C2", Weight: 1})
	if plan.TotalWeight() != 1 {
		t.Fatalf("plan total wrong")
	}

	route := logistics.Route{Name: "R", Distance: 10}
	if tm, err := route.TravelTime(5); err != nil || tm != 2 {
		t.Fatalf("travel time failed")
	}
	route.AddStop("Depot")

	vType := logistics.VehicleType{Name: "Truck", Capacity: 5}
	if !vType.CanCarry(4) {
		t.Fatalf("capacity check failed")
	}
	vType.ToggleElectric()

	weightMap := logistics.WeightMap{}
	weightMap.Register("Box", 2)
	weightMap.Register("Crate", 3)
	if weightMap.TotalWeight() != 5 {
		t.Fatalf("weight map wrong")
	}

	warehouse := logistics.Warehouse{Name: "Extra", Country: marketing.Country{Name: "IT"}, Capacity: 5}
	if err := warehouse.Store(logistics.Cargo{ID: "C3", Weight: 2}); err != nil {
		t.Fatalf("store unexpected error: %v", err)
	}
	if warehouse.Utilization() <= 0 {
		t.Fatalf("warehouse utilization zero")
	}

	carrier := logistics.Carrier{Name: "Fast"}
	carrier.AssignVehicle(logistics.Vehicle{})
	carrier.Schedule(logistics.Shipment{})
}

func TestLogisticsCenter(t *testing.T) {
	country := marketing.Country{Name: "USA"}
	route := logistics.Route{Name: "Extra", Distance: 5}

	prod := logistics.Product{Name: "Calzone"}
	cargo := logistics.Cargo{ID: "CX", Weight: 2}
	cargo.AssignProduct(&prod)

	bufferWarehouse := logistics.Warehouse{Name: "Buffer", Country: country, Capacity: 5}
	if err := bufferWarehouse.Store(cargo); err != nil {
		t.Fatalf("unexpected warehouse store error: %v", err)
	}

	plan := logistics.CargoSortingPlan{Name: "Extra", Warehouse: &bufferWarehouse}
	plan.AddCargo(cargo)

	center := logistics.LogisticsCenter{Name: "Hub", Country: country}
	center.AddRoute(route)
	center.PlanCargo(plan)

	vehicle := logistics.Vehicle{Plate: "V2", Type: logistics.VehicleType{Name: "Truck", Capacity: 10}}
	if err := vehicle.Load(cargo); err != nil {
		t.Fatalf("vehicle load error: %v", err)
	}
	if vehicle.LoadCount() != 1 {
		t.Fatalf("load count mismatch")
	}

	shipment := logistics.Shipment{ID: "S2", Route: route, Vehicle: vehicle}
	shipment.AddCargo(cargo)

	dept := logistics.LogisticsDepartment{Country: country}
	if err := dept.Dispatch(shipment); err != nil {
		t.Fatalf("dispatch failed: %v", err)
	}
}
