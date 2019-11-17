package model

// Specifies that the category of a product may be indicated by a code or by free text.
type ProductCategory1Choice struct {

	// Specifies the type of product category.
	StructuredProductCategory *ProductCategory1 `xml:"StrdPdctCtgy"`

	// Specifies the type of product category not present in the code list.
	OtherProductCategory *GenericIdentification4 `xml:"OthrPdctCtgy"`
}

func (p *ProductCategory1Choice) AddStructuredProductCategory() *ProductCategory1 {
	p.StructuredProductCategory = new(ProductCategory1)
	return p.StructuredProductCategory
}

func (p *ProductCategory1Choice) AddOtherProductCategory() *GenericIdentification4 {
	p.OtherProductCategory = new(GenericIdentification4)
	return p.OtherProductCategory
}
