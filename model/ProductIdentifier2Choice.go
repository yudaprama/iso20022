package model

// Identifies a product in coded form or free text.
type ProductIdentifier2Choice struct {

	// Specifies the type of product identifier.
	StructuredProductIdentifier *ProductIdentifier2 `xml:"StrdPdctIdr"`

	// Specifies the type of product identifier not present in the code list.
	OtherProductIdentifier *GenericIdentification4 `xml:"OthrPdctIdr"`
}

func (p *ProductIdentifier2Choice) AddStructuredProductIdentifier() *ProductIdentifier2 {
	p.StructuredProductIdentifier = new(ProductIdentifier2)
	return p.StructuredProductIdentifier
}

func (p *ProductIdentifier2Choice) AddOtherProductIdentifier() *GenericIdentification4 {
	p.OtherProductIdentifier = new(GenericIdentification4)
	return p.OtherProductIdentifier
}
