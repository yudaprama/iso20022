package model

// Product to purchase.
type Product4 struct {

	// Product code.
	ProductCode *Max70Text `xml:"PdctCd"`

	// Additional product code related to the product.
	AdditionalProductCode *Max70Text `xml:"AddtlPdctCd,omitempty"`
}

func (p *Product4) SetProductCode(value string) {
	p.ProductCode = (*Max70Text)(&value)
}

func (p *Product4) SetAdditionalProductCode(value string) {
	p.AdditionalProductCode = (*Max70Text)(&value)
}
