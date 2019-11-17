package model

// Structured information supplied to enable the matching, ie, reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an accounts receivable system.
type StructuredRemittanceInformation2 struct {

	// Specifies the nature of the referred document/transaction, eg, invoice or credit note.
	ReferredDocumentType *DocumentType1Code `xml:"RfrdDocTp,omitempty"`

	// Date associated with the referred document, eg, date of issue.
	ReferredDocumentRelatedDate *ISODate `xml:"RfrdDocRltdDt,omitempty"`

	// Amount of money and currency of a document referred to in the remittance section. The amount is typically either the original amount due and payable, or the amount actually remitted for the referred document.
	ReferredDocumentAmount []*ReferredDocumentAmount1Choice `xml:"RfrdDocAmt,omitempty"`

	// Unique and unambiguous identification of a document that distinguishes that document from another document referred to in the remittance information, usually assigned by the originator of the referred document/transaction.
	DocumentReferenceNumber *Max35Text `xml:"DocRefNb,omitempty"`

	// Unique and unambiguous reference assigned by the creditor to refer to the payment transaction.
	//
	// Usage: if available, the initiating party should provide this reference in the structured remittance information, to enable reconciliation by the creditor upon receipt of the cash.
	//
	// If the business context requires the use of a creditor reference or a payment remit identification, and only one identifier can be passed through the end-to-end chain, the creditor's reference or payment remittance identification should be quoted in the end-to-end transaction identification.
	CreditorReference *Max35Text `xml:"CdtrRef,omitempty"`

	// Identification of the organization issuing the invoice when different than the creditor or final party
	Invoicer *PartyIdentification1 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when different than the originator or debtor.
	Invoicee *PartyIdentification1 `xml:"Invcee,omitempty"`
}

func (s *StructuredRemittanceInformation2) SetReferredDocumentType(value string) {
	s.ReferredDocumentType = (*DocumentType1Code)(&value)
}

func (s *StructuredRemittanceInformation2) SetReferredDocumentRelatedDate(value string) {
	s.ReferredDocumentRelatedDate = (*ISODate)(&value)
}

func (s *StructuredRemittanceInformation2) AddReferredDocumentAmount() *ReferredDocumentAmount1Choice {
	newValue := new(ReferredDocumentAmount1Choice)
	s.ReferredDocumentAmount = append(s.ReferredDocumentAmount, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation2) SetDocumentReferenceNumber(value string) {
	s.DocumentReferenceNumber = (*Max35Text)(&value)
}

func (s *StructuredRemittanceInformation2) SetCreditorReference(value string) {
	s.CreditorReference = (*Max35Text)(&value)
}

func (s *StructuredRemittanceInformation2) AddInvoicer() *PartyIdentification1 {
	s.Invoicer = new(PartyIdentification1)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation2) AddInvoicee() *PartyIdentification1 {
	s.Invoicee = new(PartyIdentification1)
	return s.Invoicee
}
