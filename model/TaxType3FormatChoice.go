package model

// Choice of formats to express the type of taxes.
type TaxType3FormatChoice struct {

	// Standard code to specify the type of taxes.
	Code *TaxType3Code `xml:"Cd"`

	// Proprietary code to express the type of taxes.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (t *TaxType3FormatChoice) SetCode(value string) {
	t.Code = (*TaxType3Code)(&value)
}

func (t *TaxType3FormatChoice) AddProprietary() *GenericIdentification13 {
	t.Proprietary = new(GenericIdentification13)
	return t.Proprietary
}
