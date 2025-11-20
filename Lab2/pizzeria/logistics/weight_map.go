package logistics

type WeightMap struct {
	Items map[string]float64
	Limit float64
}

func (w *WeightMap) Register(name string, weight float64) {
	if w.Items == nil {
		w.Items = map[string]float64{}
	}
	w.Items[name] = weight
}

func (w WeightMap) TotalWeight() float64 {
	total := 0.0
	for _, weight := range w.Items {
		total += weight
	}
	return total
}
