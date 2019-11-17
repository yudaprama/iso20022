package model

// Execution of a redemption order.
type RedemptionMultipleExecution3 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson12 `xml:"BnfcryDtls,omitempty"`

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

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution6 `xml:"IndvExctnDtls"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the multiple order.
	TotalSettlementAmount *ActiveCurrencyAndAmount `xml:"TtlSttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction22 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionMultipleExecution3) SetMasterReference(value string) {
	r.MasterReference = (*Max35Text)(&value)
}

func (r *RedemptionMultipleExecution3) AddBeneficiaryDetails() *IndividualPerson12 {
	r.BeneficiaryDetails = new(IndividualPerson12)
	return r.BeneficiaryDetails
}

func (r *RedemptionMultipleExecution3) AddPlaceOfTrade() *PlaceOfTradeIdentification1Choice {
	r.PlaceOfTrade = new(PlaceOfTradeIdentification1Choice)
	return r.PlaceOfTrade
}

func (r *RedemptionMultipleExecution3) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleExecution3) SetRequestedFutureTradeDate(value string) {
	r.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleExecution3) SetCancellationRight(value string) {
	r.CancellationRight = (*CancellationRight1Code)(&value)
}

func (r *RedemptionMultipleExecution3) SetExtendedCancellationRight(value string) {
	r.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (r *RedemptionMultipleExecution3) AddInvestmentAccountDetails() *InvestmentAccount21 {
	r.InvestmentAccountDetails = new(InvestmentAccount21)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleExecution3) AddIndividualExecutionDetails() *RedemptionExecution6 {
	newValue := new(RedemptionExecution6)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleExecution3) SetTotalSettlementAmount(value, currency string) {
	r.TotalSettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (r *RedemptionMultipleExecution3) SetCashSettlementDate(value string) {
	r.CashSettlementDate = (*ISODate)(&value)
}

func (r *RedemptionMultipleExecution3) AddBulkCashSettlementDetails() *PaymentTransaction22 {
	r.BulkCashSettlementDetails = new(PaymentTransaction22)
	return r.BulkCashSettlementDetails
}
