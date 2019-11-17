package model

// Choice of formats to  express the type of price rate.
type PriceRateType3FormatChoice struct {

	// Standard code to  specify the type of price rate.
	Code *PriceRateType3Code `xml:"Cd"`

	// Proprietary code to  express the type of price rate.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (p *PriceRateType3FormatChoice) SetCode(value string) {
	p.Code = (*PriceRateType3Code)(&value)
}

func (p *PriceRateType3FormatChoice) AddProprietary() *GenericIdentification13 {
	p.Proprietary = new(GenericIdentification13)
	return p.Proprietary
}
