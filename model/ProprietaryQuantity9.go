package model

// Proprietary quantity format.
type ProprietaryQuantity9 struct {

	// Provides the proprietary quantity with a decimal number.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity9) SetQuantity(value string) {
	p.Quantity = (*RestrictedFINDecimalNumber)(&value)
}

func (p *ProprietaryQuantity9) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity9) SetIssuer(value string) {
	p.Issuer = (*Max4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity9) SetSchemeName(value string) {
	p.SchemeName = (*Max4AlphaNumericText)(&value)
}
