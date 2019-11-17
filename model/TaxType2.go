package model

// Specification of the tax type.
type TaxType2 struct {

	// Structured format.
	Structured *TaxType5Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxType2) SetStructured(value string) {
	t.Structured = (*TaxType5Code)(&value)
}

func (t *TaxType2) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
