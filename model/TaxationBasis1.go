package model

// Specification of the taxation basis.
type TaxationBasis1 struct {

	// Structured format.
	Structured *TaxationBasis3Code `xml:"Strd"`

	// Additional information about the type of taxation basis.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxationBasis1) SetStructured(value string) {
	t.Structured = (*TaxationBasis3Code)(&value)
}

func (t *TaxationBasis1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
