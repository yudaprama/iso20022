package model

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation7 struct {

	// Set of elements used to identify the documents referred to in the remittance information.
	ReferredDocumentInformation []*ReferredDocumentInformation3 `xml:"RfrdDocInf,omitempty"`

	// Set of elements used to provide details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount1 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification32 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification32 `xml:"Invcee,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation7) AddReferredDocumentInformation() *ReferredDocumentInformation3 {
	newValue := new(ReferredDocumentInformation3)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation7) AddReferredDocumentAmount() *RemittanceAmount1 {
	s.ReferredDocumentAmount = new(RemittanceAmount1)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation7) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation7) AddInvoicer() *PartyIdentification32 {
	s.Invoicer = new(PartyIdentification32)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation7) AddInvoicee() *PartyIdentification32 {
	s.Invoicee = new(PartyIdentification32)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation7) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}
