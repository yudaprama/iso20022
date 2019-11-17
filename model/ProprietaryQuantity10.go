package model

// Proprietary quantity format.
type ProprietaryQuantity10 struct {

	// Sign of the quantity of security.
	ShortLongPosition *ShortLong1Code `xml:"ShrtLngPos,omitempty"`

	// Provides the proprietary quantity with a decimal number.
	Quantity *RestrictedFINDecimalNumber `xml:"Qty"`

	// Identifies the type of proprietary quantity reported.
	QuantityType *Exact4AlphaNumericText `xml:"QtyTp"`

	// Provides information related to issuer in free format.
	Issuer *Max4AlphaNumericText `xml:"Issr"`

	// Name of the identification scheme.
	SchemeName *Max4AlphaNumericText `xml:"SchmeNm,omitempty"`
}

func (p *ProprietaryQuantity10) SetShortLongPosition(value string) {
	p.ShortLongPosition = (*ShortLong1Code)(&value)
}

func (p *ProprietaryQuantity10) SetQuantity(value string) {
	p.Quantity = (*RestrictedFINDecimalNumber)(&value)
}

func (p *ProprietaryQuantity10) SetQuantityType(value string) {
	p.QuantityType = (*Exact4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity10) SetIssuer(value string) {
	p.Issuer = (*Max4AlphaNumericText)(&value)
}

func (p *ProprietaryQuantity10) SetSchemeName(value string) {
	p.SchemeName = (*Max4AlphaNumericText)(&value)
}
