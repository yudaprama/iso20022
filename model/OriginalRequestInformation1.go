package model

// Set of characteristics that unambiguously identify the original global invoice financing request.
type OriginalRequestInformation1 struct {

	// Unique and unambiguous identifier of the original request message as assigned by the original sending party.
	Identification *Max35Text `xml:"Id"`

	// Date and time at which the original request message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`

	// Party that requests the invoice financing, on behalf of a creditor, as indicated in the original request message.
	FinancingRequestor *PartyIdentificationAndAccount6 `xml:"FincgRqstr,omitempty"`

	// Financial institution that receives the request from the financing requestor and forwards it to the first agent for execution, as indicated in the original request message.
	IntermediaryAgent *FinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty"`

	// Financial institution of financing requestor to which an invoice financing request is addressed, as indicated in the original request message.
	FirstAgent *FinancialInstitutionIdentification6 `xml:"FrstAgt,omitempty"`

	// Information about the validation status of the request message.
	ValidationStatusInformation *ValidationStatusInformation1 `xml:"VldtnStsInf"`

	// Information on the business status of the cancellation.
	CancellationStatusInformation *CancellationStatusInformation1 `xml:"CxlStsInf,omitempty"`
}

func (o *OriginalRequestInformation1) SetIdentification(value string) {
	o.Identification = (*Max35Text)(&value)
}

func (o *OriginalRequestInformation1) SetCreationDateTime(value string) {
	o.CreationDateTime = (*ISODateTime)(&value)
}

func (o *OriginalRequestInformation1) AddFinancingRequestor() *PartyIdentificationAndAccount6 {
	o.FinancingRequestor = new(PartyIdentificationAndAccount6)
	return o.FinancingRequestor
}

func (o *OriginalRequestInformation1) AddIntermediaryAgent() *FinancialInstitutionIdentification6 {
	o.IntermediaryAgent = new(FinancialInstitutionIdentification6)
	return o.IntermediaryAgent
}

func (o *OriginalRequestInformation1) AddFirstAgent() *FinancialInstitutionIdentification6 {
	o.FirstAgent = new(FinancialInstitutionIdentification6)
	return o.FirstAgent
}

func (o *OriginalRequestInformation1) AddValidationStatusInformation() *ValidationStatusInformation1 {
	o.ValidationStatusInformation = new(ValidationStatusInformation1)
	return o.ValidationStatusInformation
}

func (o *OriginalRequestInformation1) AddCancellationStatusInformation() *CancellationStatusInformation1 {
	o.CancellationStatusInformation = new(CancellationStatusInformation1)
	return o.CancellationStatusInformation
}
