package model

// Choice between specification of the tax exemption reason in structured or free text format.
type TaxExemptionReasonFormatChoice struct {

	// Free text form.
	Unstructured *Max140Text `xml:"Ustrd"`

	// Structured format.
	Structured *TaxExemptReason1Code `xml:"Strd"`
}

func (t *TaxExemptionReasonFormatChoice) SetUnstructured(value string) {
	t.Unstructured = (*Max140Text)(&value)
}

func (t *TaxExemptionReasonFormatChoice) SetStructured(value string) {
	t.Structured = (*TaxExemptReason1Code)(&value)
}
