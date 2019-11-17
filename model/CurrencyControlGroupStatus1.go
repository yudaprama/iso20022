package model

// Provides the currency control group status details.
type CurrencyControlGroupStatus1 struct {

	// Original underlying message references for which the status advice is provided.
	OriginalReferences *OriginalMessage3 `xml:"OrgnlRefs"`

	// Party registering the currency control contract.
	ReportingParty *TradeParty2 `xml:"RptgPty"`

	// Agent which registers the currency control contract.
	RegistrationAgent *BranchAndFinancialInstitutionIdentification5 `xml:"RegnAgt"`

	// For daily reporting this is the day to which the transaction data in the status message refers to.
	// For periodic reporting this is the first and the last day to which the transaction data in the status message refers.
	ReportingPeriod *Period4Choice `xml:"RptgPrd,omitempty"`

	// Provides the status for the full report.
	Status *StatisticalReportingStatus1Code `xml:"Sts,omitempty"`

	// Provides detailed information on the status reason.
	StatusReason []*ValidationStatusReason1 `xml:"StsRsn,omitempty"`

	// Provides the date and time when the status was issued.
	StatusDateTime *ISODateTime `xml:"StsDtTm,omitempty"`
}

func (c *CurrencyControlGroupStatus1) AddOriginalReferences() *OriginalMessage3 {
	c.OriginalReferences = new(OriginalMessage3)
	return c.OriginalReferences
}

func (c *CurrencyControlGroupStatus1) AddReportingParty() *TradeParty2 {
	c.ReportingParty = new(TradeParty2)
	return c.ReportingParty
}

func (c *CurrencyControlGroupStatus1) AddRegistrationAgent() *BranchAndFinancialInstitutionIdentification5 {
	c.RegistrationAgent = new(BranchAndFinancialInstitutionIdentification5)
	return c.RegistrationAgent
}

func (c *CurrencyControlGroupStatus1) AddReportingPeriod() *Period4Choice {
	c.ReportingPeriod = new(Period4Choice)
	return c.ReportingPeriod
}

func (c *CurrencyControlGroupStatus1) SetStatus(value string) {
	c.Status = (*StatisticalReportingStatus1Code)(&value)
}

func (c *CurrencyControlGroupStatus1) AddStatusReason() *ValidationStatusReason1 {
	newValue := new(ValidationStatusReason1)
	c.StatusReason = append(c.StatusReason, newValue)
	return newValue
}

func (c *CurrencyControlGroupStatus1) SetStatusDateTime(value string) {
	c.StatusDateTime = (*ISODateTime)(&value)
}
