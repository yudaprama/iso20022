package model

// Choice of formats for taxable income per share calculated.
type TaxableIncomePerShareCalculated2Choice struct {

	// Taxable income per share (TIS) calculated expressed as a code.
	Code *TaxableIncomePerShareCalculated2Code `xml:"Cd"`

	// Taxable income per share calculated (TIS) expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxableIncomePerShareCalculated2Choice) SetCode(value string) {
	t.Code = (*TaxableIncomePerShareCalculated2Code)(&value)
}

func (t *TaxableIncomePerShareCalculated2Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
