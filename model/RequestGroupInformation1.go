package model

// Set of characteristics that unambiguously identify the global invoice financing request.
type RequestGroupInformation1 struct {

	// Point to point reference assigned by the financing requestor  to unambiguously identify the invoice financing request message.
	//
	// Usage: The financing requestor has to make sure that 'GroupIdentification' is unique for a pre-agreed period.
	GroupIdentification *Max35Text `xml:"GrpId"`

	// Date and time on which the invoice financing request was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// User identification or any user key that allows to check if the financing requestor is allowed to ask for invoice financing.
	//
	// Usage: the content is not of a technical nature, but reflects the organisational structure at the requesting side.
	// The authorisation element can typically be used in case the financing requestor acts on behalf of one or more suppliers.
	Authorisation []*Max128Text `xml:"Authstn,omitempty"`

	// Specifies the number of single invoice financing requests included in the bulk request message.
	NumberOfInvoiceRequests *Max15NumericText `xml:"NbOfInvcReqs,omitempty"`

	// Total amount of the bulk invoice financing request. It is composed by the sum of the total amounts of all invoices included in the financing request.
	TotalBulkInvoiceAmount *ActiveCurrencyAndAmount `xml:"TtlBlkInvcAmt,omitempty"`

	// Reference currency of the invoice financing request.
	Currency *CurrencyCode `xml:"Ccy"`

	// Specifies the financing method related to invoice financing (eg collection mandate).
	FinancingAgreement *Max350Text `xml:"FincgAgrmt,omitempty"`

	// Party that requests the invoice financing, on behalf of a creditor.
	FinancingRequestor *PartyIdentificationAndAccount6 `xml:"FincgRqstr"`

	// Financial institution that receives the request from the financing requestor and forwards it to the first agent for execution.
	IntermediaryAgent *FinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty"`

	// Financial institution of financing requestor to which an invoice financing request is addressed.
	FirstAgent *FinancialInstitutionIdentification6 `xml:"FrstAgt"`

	// Agreements between financing requestor and his bank concerning conditions about the service of invoice financing, based on specific contractual schemes.
	AgreementClauses []*AgreementClauses1 `xml:"AgrmtClauses,omitempty"`

	// Additional information about the financing request.
	AdditionalInformation []*AdditionalInformation1 `xml:"AddtlInf,omitempty"`
}

func (r *RequestGroupInformation1) SetGroupIdentification(value string) {
	r.GroupIdentification = (*Max35Text)(&value)
}

func (r *RequestGroupInformation1) SetCreationDateTime(value string) {
	r.CreationDateTime = (*ISODateTime)(&value)
}

func (r *RequestGroupInformation1) AddAuthorisation(value string) {
	r.Authorisation = append(r.Authorisation, (*Max128Text)(&value))
}

func (r *RequestGroupInformation1) SetNumberOfInvoiceRequests(value string) {
	r.NumberOfInvoiceRequests = (*Max15NumericText)(&value)
}

func (r *RequestGroupInformation1) SetTotalBulkInvoiceAmount(value, currency string) {
	r.TotalBulkInvoiceAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RequestGroupInformation1) SetCurrency(value string) {
	r.Currency = (*CurrencyCode)(&value)
}

func (r *RequestGroupInformation1) SetFinancingAgreement(value string) {
	r.FinancingAgreement = (*Max350Text)(&value)
}

func (r *RequestGroupInformation1) AddFinancingRequestor() *PartyIdentificationAndAccount6 {
	r.FinancingRequestor = new(PartyIdentificationAndAccount6)
	return r.FinancingRequestor
}

func (r *RequestGroupInformation1) AddIntermediaryAgent() *FinancialInstitutionIdentification6 {
	r.IntermediaryAgent = new(FinancialInstitutionIdentification6)
	return r.IntermediaryAgent
}

func (r *RequestGroupInformation1) AddFirstAgent() *FinancialInstitutionIdentification6 {
	r.FirstAgent = new(FinancialInstitutionIdentification6)
	return r.FirstAgent
}

func (r *RequestGroupInformation1) AddAgreementClauses() *AgreementClauses1 {
	newValue := new(AgreementClauses1)
	r.AgreementClauses = append(r.AgreementClauses, newValue)
	return newValue
}

func (r *RequestGroupInformation1) AddAdditionalInformation() *AdditionalInformation1 {
	newValue := new(AdditionalInformation1)
	r.AdditionalInformation = append(r.AdditionalInformation, newValue)
	return newValue
}
