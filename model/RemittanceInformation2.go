package model

// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle.
type RemittanceInformation2 struct {

	// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`
}

func (r *RemittanceInformation2) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}
