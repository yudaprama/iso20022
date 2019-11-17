package model

// Proprietary quantity format.
type ProprietaryQuantity7 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos,omitempty"`

	// Provides the proprietary quantity with a decimal number.
	Quantity *DecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max35Text `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max35Text `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity7) SetShortLongPosition(value string) {
	p.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (p *ProprietaryQuantity7) SetQuantity(value string) {
	p.Quantity = (*DecimalNumber)(&value)
}

func (p *ProprietaryQuantity7) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity7) SetIssuer(value string) {
	p.Issuer = (*Max35Text)(&value)
}

func (p *ProprietaryQuantity7) SetSchemeName(value string) {
	p.SchemeName = (*Max35Text)(&value)
}
