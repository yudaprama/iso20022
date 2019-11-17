package model

// Execution of a redemption order.
type RedemptionBulkExecution2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Execution of a redemption order.
	IndividualExecutionDetails []*RedemptionExecution3 `xml:"IndvExctnDtls"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Payment transaction related to the execution of an investment fund transaction.
	BulkCashSettlementDetails *PaymentTransaction18 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionBulkExecution2) SetPlaceOfTrade(value string) {
	r.PlaceOfTrade = (*CountryCode)(&value)
}

func (r *RedemptionBulkExecution2) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionBulkExecution2) AddCancellationRight() *CancellationRight1 {
	r.CancellationRight = new(CancellationRight1)
	return r.CancellationRight
}

func (r *RedemptionBulkExecution2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	r.FinancialInstrumentDetails = new(FinancialInstrument6)
	return r.FinancialInstrumentDetails
}

func (r *RedemptionBulkExecution2) AddIndividualExecutionDetails() *RedemptionExecution3 {
	newValue := new(RedemptionExecution3)
	r.IndividualExecutionDetails = append(r.IndividualExecutionDetails, newValue)
	return newValue
}

func (r *RedemptionBulkExecution2) SetRequestedSettlementCurrency(value string) {
	r.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionBulkExecution2) SetRequestedNAVCurrency(value string) {
	r.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (r *RedemptionBulkExecution2) AddBulkCashSettlementDetails() *PaymentTransaction18 {
	r.BulkCashSettlementDetails = new(PaymentTransaction18)
	return r.BulkCashSettlementDetails
}
