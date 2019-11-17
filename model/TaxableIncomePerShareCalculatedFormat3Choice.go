package model

// Choice between a standard code or proprietary code to specify whether the taxable income per share or per dividend is calculated.
type TaxableIncomePerShareCalculatedFormat3Choice struct {

	// Standard code to specify whether the fund calculates the taxable income per dividend/taxable income per share (TID/TIS).
	Code *CorporateActionTaxableIncomePerShareCalculated1Code `xml:"Cd"`

	// Proprietary identification to specify whether the fund calculates the taxable income per dividend/taxable income per share (TID/TIS).
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (t *TaxableIncomePerShareCalculatedFormat3Choice) SetCode(value string) {
	t.Code = (*CorporateActionTaxableIncomePerShareCalculated1Code)(&value)
}

func (t *TaxableIncomePerShareCalculatedFormat3Choice) AddProprietary() *GenericIdentification30 {
	t.Proprietary = new(GenericIdentification30)
	return t.Proprietary
}
