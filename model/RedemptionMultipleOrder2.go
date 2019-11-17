package model

// Instruction from an investor to sell investment fund units back to the fund.
type RedemptionMultipleOrder2 struct {

	// Market in which the advised trade transaction was executed.
	PlaceOfTrade *CountryCode `xml:"PlcOfTrad,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *ISODateTime `xml:"XpryDtTm,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Account impacted by an investment fund order.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Additional information about the investor.
	BeneficiaryDetails *IndividualPerson2 `xml:"BnfcryDtls,omitempty"`

	// Instruction from an investor to sell investment fund units back to the fund.
	IndividualOrderDetails []*RedemptionOrder4 `xml:"IndvOrdrDtls"`

	// Payment transaction resulting from the investment fund order execution.
	BulkCashSettlementDetails *PaymentTransaction15 `xml:"BlkCshSttlmDtls,omitempty"`
}

func (r *RedemptionMultipleOrder2) SetPlaceOfTrade(value string) {
	r.PlaceOfTrade = (*CountryCode)(&value)
}

func (r *RedemptionMultipleOrder2) SetOrderDateTime(value string) {
	r.OrderDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleOrder2) SetExpiryDateTime(value string) {
	r.ExpiryDateTime = (*ISODateTime)(&value)
}

func (r *RedemptionMultipleOrder2) AddCancellationRight() *CancellationRight1 {
	r.CancellationRight = new(CancellationRight1)
	return r.CancellationRight
}

func (r *RedemptionMultipleOrder2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	r.InvestmentAccountDetails = new(InvestmentAccount13)
	return r.InvestmentAccountDetails
}

func (r *RedemptionMultipleOrder2) AddBeneficiaryDetails() *IndividualPerson2 {
	r.BeneficiaryDetails = new(IndividualPerson2)
	return r.BeneficiaryDetails
}

func (r *RedemptionMultipleOrder2) AddIndividualOrderDetails() *RedemptionOrder4 {
	newValue := new(RedemptionOrder4)
	r.IndividualOrderDetails = append(r.IndividualOrderDetails, newValue)
	return newValue
}

func (r *RedemptionMultipleOrder2) AddBulkCashSettlementDetails() *PaymentTransaction15 {
	r.BulkCashSettlementDetails = new(PaymentTransaction15)
	return r.BulkCashSettlementDetails
}
