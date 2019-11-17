package model

// Execution of a switch order.
type SwitchExecution3 struct {

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Amount of money used to determine the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Redemption leg of a switch order execution.
	RedemptionLegDetails []*SwitchRedemptionLegExecution2 `xml:"RedLegDtls"`

	// Subscription leg of a switch order execution.
	SubscriptionLegDetails []*SwitchSubscriptionLegExecution2 `xml:"SbcptLegDtls"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction14 `xml:"CshSttlmDtls,omitempty"`

	// Currency exchange related to the execution of an investment fund order.
	ForeignExchangeDetails []*ForeignExchangeTerms4 `xml:"FXDtls,omitempty"`
}

func (s *SwitchExecution3) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchExecution3) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SwitchExecution3) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchExecution3) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SwitchExecution3) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution3) AddRedemptionLegDetails() *SwitchRedemptionLegExecution2 {
	newValue := new(SwitchRedemptionLegExecution2)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution3) AddSubscriptionLegDetails() *SwitchSubscriptionLegExecution2 {
	newValue := new(SwitchSubscriptionLegExecution2)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution3) AddCashSettlementDetails() *PaymentTransaction14 {
	s.CashSettlementDetails = new(PaymentTransaction14)
	return s.CashSettlementDetails
}

func (s *SwitchExecution3) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}
