package model

// Transfer from one investment fund/fund class to another investment fund or investment fund class by the investor. A switch is composed of one or several subscription legs, and one or several redemption legs.
type SwitchOrder2 struct {

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls"`

	// Amount of money used to derive the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *ISODateTime `xml:"XpryDtTm,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1 `xml:"CxlRght,omitempty"`

	// Part of an investment fund switch order that is a redemption.
	RedemptionLegDetails []*SwitchRedemptionLegOrder2 `xml:"RedLegDtls"`

	// Part of an investment fund switch order that is a subscription.
	SubscriptionLegDetails []*SwitchSubscriptionLegOrder2 `xml:"SbcptLegDtls"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction20 `xml:"CshSttlmDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms5 `xml:"FXDtls,omitempty"`
}

func (s *SwitchOrder2) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder2) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrder2) AddInvestmentAccountDetails() *InvestmentAccount13 {
	s.InvestmentAccountDetails = new(InvestmentAccount13)
	return s.InvestmentAccountDetails
}

func (s *SwitchOrder2) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) SetExpiryDateTime(value string) {
	s.ExpiryDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder2) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder2) AddCancellationRight() *CancellationRight1 {
	s.CancellationRight = new(CancellationRight1)
	return s.CancellationRight
}

func (s *SwitchOrder2) AddRedemptionLegDetails() *SwitchRedemptionLegOrder2 {
	newValue := new(SwitchRedemptionLegOrder2)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder2) AddSubscriptionLegDetails() *SwitchSubscriptionLegOrder2 {
	newValue := new(SwitchSubscriptionLegOrder2)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder2) AddCashSettlementDetails() *PaymentTransaction20 {
	s.CashSettlementDetails = new(PaymentTransaction20)
	return s.CashSettlementDetails
}

func (s *SwitchOrder2) AddForeignExchangeDetails() *ForeignExchangeTerms5 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms5)
	return s.ForeignExchangeDetails
}
