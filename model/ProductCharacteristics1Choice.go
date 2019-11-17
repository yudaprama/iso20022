package model

// Specifies that the category of a product may be indicated by a code or by free text.
type ProductCharacteristics1Choice struct {

	// Specifies the type of product characteristic.
	StructuredProductCharacteristics *ProductCharacteristics1 `xml:"StrdPdctChrtcs"`

	// Specifies the type of product characteristic not present in the code list.
	OtherProductCharacteristics *GenericIdentification4 `xml:"OthrPdctChrtcs"`
}

func (p *ProductCharacteristics1Choice) AddStructuredProductCharacteristics() *ProductCharacteristics1 {
	p.StructuredProductCharacteristics = new(ProductCharacteristics1)
	return p.StructuredProductCharacteristics
}

func (p *ProductCharacteristics1Choice) AddOtherProductCharacteristics() *GenericIdentification4 {
	p.OtherProductCharacteristics = new(GenericIdentification4)
	return p.OtherProductCharacteristics
}
