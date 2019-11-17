package model

// Specification of the tax type.
type TaxType1 struct {

	// Structured format.
	Structured *TaxType7Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxType1) SetStructured(value string) {
	t.Structured = (*TaxType7Code)(&value)
}

func (t *TaxType1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
