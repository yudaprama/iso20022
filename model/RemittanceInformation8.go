package model

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system.
type RemittanceInformation8 struct {

	// Unique identification, assigned by the originator, to unambiguously identify the remittance information within the message.
	RemittanceIdentification *Max35Text `xml:"RmtId,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in an unstructured form.
	Unstructured []*Max140Text `xml:"Ustrd,omitempty"`

	// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
	Structured []*StructuredRemittanceInformation10 `xml:"Strd,omitempty"`

	// Set of elements used to provide information on the original transactions, to which the remittance message refers.
	OriginalPaymentInformation *OriginalPaymentInformation6 `xml:"OrgnlPmtInf"`
}

func (r *RemittanceInformation8) SetRemittanceIdentification(value string) {
	r.RemittanceIdentification = (*Max35Text)(&value)
}

func (r *RemittanceInformation8) AddUnstructured(value string) {
	r.Unstructured = append(r.Unstructured, (*Max140Text)(&value))
}

func (r *RemittanceInformation8) AddStructured() *StructuredRemittanceInformation10 {
	newValue := new(StructuredRemittanceInformation10)
	r.Structured = append(r.Structured, newValue)
	return newValue
}

func (r *RemittanceInformation8) AddOriginalPaymentInformation() *OriginalPaymentInformation6 {
	r.OriginalPaymentInformation = new(OriginalPaymentInformation6)
	return r.OriginalPaymentInformation
}
