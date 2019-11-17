package model

// Set of information related to the cancellation request, such as actors involved and identification of the original multiple invoice financing request to which the cancellation request is referring.
type CancellationRequestInformation1 struct {

	// Unique and unambiguous identifier of the original financing request message as assigned by the original sending party.
	OriginalGroupIdentification *Max35Text `xml:"OrgnlGrpId"`

	// Date and time at which the original financing request message was created.
	OriginalCreationDateTime *ISODateTime `xml:"OrgnlCreDtTm"`

	// Specifies the number of single invoice financing requests included in the original financing request message.
	NumberOfInvoiceRequests *Max15NumericText `xml:"NbOfInvcReqs,omitempty"`

	// Total amount of the bulk invoice financing request. It is composed by the sum of the total amounts of all invoices included in the original financing request message.
	TotalBulkInvoiceAmount *ActiveCurrencyAndAmount `xml:"TtlBlkInvcAmt,omitempty"`

	// Further details on the cancellation request information, in an uncoded form.
	CancellationReason *Max105Text `xml:"CxlRsn"`

	// Party that requests the cancellation of a financing request previously sent.
	FinancingRequestor *PartyIdentificationAndAccount6 `xml:"FincgRqstr,omitempty"`

	// Financial institution that receives the request from the financing requestor and forwards it to the first agent for execution.
	IntermediaryAgent *FinancialInstitutionIdentification6 `xml:"IntrmyAgt,omitempty"`

	// Financial institution of financing requestor to which an invoice financing cancellation request is addressed.
	FirstAgent *FinancialInstitutionIdentification6 `xml:"FrstAgt,omitempty"`
}

func (c *CancellationRequestInformation1) SetOriginalGroupIdentification(value string) {
	c.OriginalGroupIdentification = (*Max35Text)(&value)
}

func (c *CancellationRequestInformation1) SetOriginalCreationDateTime(value string) {
	c.OriginalCreationDateTime = (*ISODateTime)(&value)
}

func (c *CancellationRequestInformation1) SetNumberOfInvoiceRequests(value string) {
	c.NumberOfInvoiceRequests = (*Max15NumericText)(&value)
}

func (c *CancellationRequestInformation1) SetTotalBulkInvoiceAmount(value, currency string) {
	c.TotalBulkInvoiceAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (c *CancellationRequestInformation1) SetCancellationReason(value string) {
	c.CancellationReason = (*Max105Text)(&value)
}

func (c *CancellationRequestInformation1) AddFinancingRequestor() *PartyIdentificationAndAccount6 {
	c.FinancingRequestor = new(PartyIdentificationAndAccount6)
	return c.FinancingRequestor
}

func (c *CancellationRequestInformation1) AddIntermediaryAgent() *FinancialInstitutionIdentification6 {
	c.IntermediaryAgent = new(FinancialInstitutionIdentification6)
	return c.IntermediaryAgent
}

func (c *CancellationRequestInformation1) AddFirstAgent() *FinancialInstitutionIdentification6 {
	c.FirstAgent = new(FinancialInstitutionIdentification6)
	return c.FirstAgent
}
