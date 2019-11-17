package model

// Choice of formats to express the type of taxes.
type TaxTypeFormat1Choice struct {

	// Standard code to specify the type of taxes.
	Code *TaxType15Code `xml:"Cd"`

	// Proprietary code to express the type of taxes.
	Proprietary *GenericIdentification13 `xml:"Prtry"`
}

func (t *TaxTypeFormat1Choice) SetCode(value string) {
	t.Code = (*TaxType15Code)(&value)
}

func (t *TaxTypeFormat1Choice) AddProprietary() *GenericIdentification13 {
	t.Proprietary = new(GenericIdentification13)
	return t.Proprietary
}
