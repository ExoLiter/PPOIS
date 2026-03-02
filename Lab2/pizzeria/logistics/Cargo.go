package logistics

import "fmt"

type Cargo struct {
	ID       string
	Weight   float64
	Material *Material
	Product  *Product
}

func (c *Cargo) AssignProduct(prod *Product) {
	c.Product = prod
	if prod != nil {
		c.Weight += float64(prod.IngredientCount())
	}
}

func (c Cargo) Describe() string {
	return fmt.Sprintf("cargo:%s:%.1f", c.ID, c.Weight)
}
