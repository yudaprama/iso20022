package model

// Information supplied to enable the matching/reconciliation of an entry with the items that the payment is intended to settle, such as commercial invoices in an accounts' receivable system, in a structured form.
type StructuredRemittanceInformation13 struct {

	// Provides the identification and the content of the referred document.
	ReferredDocumentInformation []*ReferredDocumentInformation7 `xml:"RfrdDocInf,omitempty"`

	// Provides details on the amounts of the referred document.
	ReferredDocumentAmount *RemittanceAmount2 `xml:"RfrdDocAmt,omitempty"`

	// Reference information provided by the creditor to allow the identification of the underlying documents.
	CreditorReferenceInformation *CreditorReferenceInformation2 `xml:"CdtrRefInf,omitempty"`

	// Identification of the organisation issuing the invoice, when it is different from the creditor or ultimate creditor.
	Invoicer *PartyIdentification43 `xml:"Invcr,omitempty"`

	// Identification of the party to whom an invoice is issued, when it is different from the debtor or ultimate debtor.
	Invoicee *PartyIdentification43 `xml:"Invcee,omitempty"`

	// Provides remittance information about a payment made for tax-related purposes.
	TaxRemittance *TaxInformation4 `xml:"TaxRmt,omitempty"`

	// Provides remittance information about a payment for garnishment-related purposes.
	GarnishmentRemittance *Garnishment1 `xml:"GrnshmtRmt,omitempty"`

	// Additional information, in free text form, to complement the structured remittance information.
	AdditionalRemittanceInformation []*Max140Text `xml:"AddtlRmtInf,omitempty"`
}

func (s *StructuredRemittanceInformation13) AddReferredDocumentInformation() *ReferredDocumentInformation7 {
	newValue := new(ReferredDocumentInformation7)
	s.ReferredDocumentInformation = append(s.ReferredDocumentInformation, newValue)
	return newValue
}

func (s *StructuredRemittanceInformation13) AddReferredDocumentAmount() *RemittanceAmount2 {
	s.ReferredDocumentAmount = new(RemittanceAmount2)
	return s.ReferredDocumentAmount
}

func (s *StructuredRemittanceInformation13) AddCreditorReferenceInformation() *CreditorReferenceInformation2 {
	s.CreditorReferenceInformation = new(CreditorReferenceInformation2)
	return s.CreditorReferenceInformation
}

func (s *StructuredRemittanceInformation13) AddInvoicer() *PartyIdentification43 {
	s.Invoicer = new(PartyIdentification43)
	return s.Invoicer
}

func (s *StructuredRemittanceInformation13) AddInvoicee() *PartyIdentification43 {
	s.Invoicee = new(PartyIdentification43)
	return s.Invoicee
}

func (s *StructuredRemittanceInformation13) AddTaxRemittance() *TaxInformation4 {
	s.TaxRemittance = new(TaxInformation4)
	return s.TaxRemittance
}

func (s *StructuredRemittanceInformation13) AddGarnishmentRemittance() *Garnishment1 {
	s.GarnishmentRemittance = new(Garnishment1)
	return s.GarnishmentRemittance
}

func (s *StructuredRemittanceInformation13) AddAdditionalRemittanceInformation(value string) {
	s.AdditionalRemittanceInformation = append(s.AdditionalRemittanceInformation, (*Max140Text)(&value))
}
