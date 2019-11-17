package model

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation11 struct {

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation13 `xml:"Strd,omitempty"`
}

func (r *RemittanceInformation11) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation11) AddStructured() *StructuredRemittanceInformation13 {
	newValue := new(StructuredRemittanceInformation13)
	r.Structured = append(r.Structured, newValue)
	return newValue
}
