package model

// Information supplied to enable the matching of an entry with the items that the transfer is intended to settle, eg, commercial invoices in an accounts' receivable system.
type RemittanceInformation3Choice struct {

	// Information, in free text form, to enable the matching, ie reconciliation, (reconciliation) of a payment with the items that the payment is intended to settle, such as commercial invoices in an accounts receivable system.
	Unstructured *Max140Text `xml:"Ustrd"`

	// Information in structured form, that is supplied to enable the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an account receivable system.
	Structured *StructuredRemittanceInformation2 `xml:"Strd"`
}

func (r *RemittanceInformation3Choice) SetUnstructured(value string) {
	r.Unstructured = (*Max140Text)(&value)
}

func (r *RemittanceInformation3Choice) AddStructured() *StructuredRemittanceInformation2 {
	r.Structured = new(StructuredRemittanceInformation2)
	return r.Structured
}
