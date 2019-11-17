package model

// Choice of formats to express the amount price type.
type AmountPriceType1FormatChoice struct {

	// Standard code to specify the amount price type.
	Code *AmountPriceType1Code `xml:"Cd"`

	// Proprietary code to express the amount price type.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (a *AmountPriceType1FormatChoice) SetCode(value string) {
	a.Code = (*AmountPriceType1Code)(&value)
}

func (a *AmountPriceType1FormatChoice) AddProprietary() *GenericIdentification13 {
	a.Proprietary = new(GenericIdentification13)
	return a.Proprietary
}
