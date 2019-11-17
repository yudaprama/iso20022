package model

// Structured information supplied to enable the matching, i.e.  reconciliation, of a payment with the items that the payment is intended to settle, eg, commercial invoices in an accounts receivable system.
type StructuredRemittanceInformation6 struct {

	// Reference information to allow the identification of the underlying reference documents.
	ReferredDocumentInformation *ReferredDocumentInformation1 `xml:"RfrdDocInf,omitempty"`

	// Date associated with the referred document, eg, date of issue.
	ReferredDocumentRelatedDate *ISODate `xml:"RfrdDocRltdDt,omitempty"`

	// Amount of money and currency of a document referred to in the remittance section. The amount is typically either the original amount due and payable, or the amount actually remitted for the referred document.
	ReferredDocumentAmount []*ReferredDocumentAmount1Choice `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation1 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organization issuing the invoice when different than the creditor or final party.
	Invoicer *PartyIdentification8 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when different than the originator or debtor.
	Invoicee *PartyIdentification8 `xml:"Invcee,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation *Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation6) AddReferredDocumentInformation() *ReferredDocumentInformation1 {
	s.ReferredDocumentInformation = new(ReferredDocumentInformation1)
	return s.ReferredDocumentInformation
}

func (s *StructuredRemittanceInformation6) SetReferredDocumentRelatedDate(value string) {
	s.ReferredDocumentRelatedDate = (*ISODate)(&value)
}

func (s *StructuredRemittanceInformation6) AddReferredDocumentAmount() *ReferredDocumentAmount1Choice {
	newValue := new(ReferredDocumentAmount1Choice)
	s.ReferredDocumentAmount = append(s.ReferredDocumentAmount, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation6) AddCreditorReferenceInformation() *CreditorReferenceInformation1 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation1)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation6) AddInvoicer() *PartyIdentification8 {
	s.Invoicer = new(PartyIdentification8)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation6) AddInvoicee() *PartyIdentification8 {
	s.Invoicee = new(PartyIdentification8)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation6) SetAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = (*Max140Text)(&value)
}
