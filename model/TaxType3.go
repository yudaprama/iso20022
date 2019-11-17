package model

// Specification of the tax type.
type TaxType3 struct {

	// Structured format.
	Structured *TaxType6Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxType3) SetStructured(value string) {
	t.Structured = (*TaxType6Code)(&value)
}

func (t *TaxType3) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
