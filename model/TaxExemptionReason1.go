package model

// Specification of the tax exemption reason.
type TaxExemptionReason1 struct {

	// Structured format.
	Structured *TaxExemptReason2Code `xml:"Strd"`

	// Additional information about the type of tax.
	AdditionalInformation *Max350Text `xml:"AddtlInf,omitempty"`
}

func (t *TaxExemptionReason1) SetStructured(value string) {
	t.Structured = (*TaxExemptReason2Code)(&value)
}

func (t *TaxExemptionReason1) SetAdditionalInformation(value string) {
	t.AdditionalInformation = (*Max350Text)(&value)
}
