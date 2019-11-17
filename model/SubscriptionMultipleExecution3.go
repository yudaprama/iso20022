package model

// Execution of a subscription order.
type SubscriptionMultipleExecution3 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *PlaceOfTradeIdentification1Choice `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

	// Execution of a subscription order.
	IndividualExecutionDetails []*SubscriptionExecution6 `xml:"IndvExctnDtls"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction24 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (s *SubscriptionMultipleExecution3) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SubscriptionMultipleExecution3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	s.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return s.PlaceOfTrade
}

func (s *SubscriptionMultipleExecution3) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SubscriptionMultipleExecution3) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SubscriptionMultipleExecution3) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SubscriptionMultipleExecution3) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SubscriptionMultipleExecution3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SubscriptionMultipleExecution3) AddBeneficiaryDetails() *IndividualPerson12 {
	s.BeneficiaryDetails = new(IndividualPerson12)
	return s.BeneficiaryDetails
}

func (s *SubscriptionMultipleExecution3) AddIndividualExecutionDetails() *SubscriptionExecution6 {
	newValue := new(SubscriptionExecution6)
	s.IndividualExecutionDetails = append(s.IndividualExecutionDetails, newValue)
	return newValue
}

func (s *SubscriptionMultipleExecution3) SetTotalSettlementAmount(value, currency string) {
	s.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionMultipleExecution3) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SubscriptionMultipleExecution3) AddBulkCashSettlementDetails() *PaymentTransaction24 {
	s.BulkCashSettlementDetails = new(PaymentTransaction24)
	return s.BulkCashSettlementDetails
}
