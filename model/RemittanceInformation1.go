package model

// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system.
type RemittanceInformation1 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system in a structured form.
	Structured []*StructuredRemittanceInformation6 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation1) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation1) AddStructured() *StructuredRemittanceInformation6 {
	newValue := new(StructuredRemittanceInformation6)
	r.Structured = append(r.Structured, newValue)
	return newValue
}
