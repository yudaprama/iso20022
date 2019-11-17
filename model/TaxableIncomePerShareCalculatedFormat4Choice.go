package model

// Choice between a standard code or proprietary code to specify whether the taxable income per share or per dividend is calculated.
type TaxableIncomePerShareCalculatedFormat4Choice struct {

	// Standard code to specify whether the fund calculates the taxable income per dividend/taxable income per share (TID/TIS).
	Code *CorporateActionTaxableIncomePerShareCalculated1Code `xml:"Cd"`

	// Proprietary identification to specify whether the fund calculates the taxable income per dividend/taxable income per share (TID/TIS).
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (t *TaxableIncomePerShareCalculatedFormat4Choice) SetCode(value string) {
	t.Code = (*CorporateActionTaxableIncomePerShareCalculated1Code)(&value)
}

func (t *TaxableIncomePerShareCalculatedFormat4Choice) AddProprietary() *GenericIdentification47 {
	t.Proprietary = new(GenericIdentification47)
	return t.Proprietary
}
