package model

// Specifies the category of the product.
type ProductCategory1 struct {

	// Specifies the type of product category by means of a code.
	Type *ProductCategory1Code `xml:"Tp"`

	// Specifies the category of a product.
	Category *Max35Text `xml:"Ctgy"`
}

func (p *ProductCategory1) SetType(value string) {
	p.Type = (*ProductCategory1Code)(&value)
}

func (p *ProductCategory1) SetCategory(value string) {
	p.Category = (*Max35Text)(&value)
}
