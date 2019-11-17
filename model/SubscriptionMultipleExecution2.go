package model

// Execution of a subscription order.
type SubscriptionMultipleExecution2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Execution of a subscription order.
	IndividualExecutionDetails []*SubscriptionExecution4 `xml:"IndvExctnDtls"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction13 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionMultipleExecution2) SetPlaceOfTrade(value string) {
	s.PlaceOfTrade = (*CountryCode)(&value)
}

func (s *SubscriptionMultipleExecution2) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionMultipleExecution2) AddCancellationRight() *CancellationRight1 {
	s.CancellationRight = new(CancellationRight1)
	return s.CancellationRight
}

func (s *SubscriptionMultipleExecution2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionMultipleExecution2) AddBeneficiaryDetails() *IndividualPerson2 {
	s.BeneficiaryDetails = new(IndividualPerson2)
	return s.BeneficiaryDetails
}

func (s *SubscriptionMultipleExecution2) AddIndividualExecutionDetails() *SubscriptionExecution4 {
	newValue := new(SubscriptionExecution4)
	s.IndividualExecutionDetails = append(s.IndividualExecutionDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleExecution2) AddBulkCashSettlementDetails() *PaymentTransaction13 {
	s.BulkCashSettlementDetails = new(PaymentTransaction13)
	return s.BulkCashSettlementDetails
}
