package model

// Specification of the taxable income per share.
type TaxableIncomePerShareCalculated1 struct {

	// Structured format.
	Structured *TaxableIncomePerShareCalculated1Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxableIncomePerShareCalculated1) SetStructured(value string) {
	t.Structured = (*TaxableIncomePerShareCalculated1Code)(&value)
}

func (t *TaxableIncomePerShareCalculated1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
