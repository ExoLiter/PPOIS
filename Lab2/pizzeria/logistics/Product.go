package logistics

type Product struct {
	Name        string
	Size        string
	Ingredients []Material
}

func (p *Product) AddIngredient(material Material) {
	p.Ingredients = append(p.Ingredients, material)
}

func (p Product) IngredientCount() int {
	return len(p.Ingredients)
}
