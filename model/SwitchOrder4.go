package model

// Transfer from one investment fund/fund class to another investment fund or investment fund class by the investor. A switch is composed of one or several subscription legs, and one or several redemption legs.
type SwitchOrder4 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveOrHistoricCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Date on which the order expires.
	ExpiryDateTime *DateAndDateTimeChoice `xml:"XpryDtTm,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveOrHistoricCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveOrHistoricCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary8 `xml:"RltdPtyDtls,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	CancellationRight *CancellationRight1Code `xml:"CxlRght,omitempty"`

	// Cancellation right of an investor with respect to an investment fund order.
	ExtendedCancellationRight *Extended350Code `xml:"XtndedCxlRght,omitempty"`

	// Part of an investment fund switch order that is a redemption.
	RedemptionLegDetails []*SwitchRedemptionLegOrder3 `xml:"RedLegDtls"`

	// Part of an investment fund switch order that is a subscription.
	SubscriptionLegDetails []*SwitchSubscriptionLegOrder3 `xml:"SbcptLegDtls"`

	// Payment processes required to transfer cash from the debtor to the creditor.
	CashSettlementDetails *PaymentTransaction25 `xml:"CshSttlmDtls,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails *ForeignExchangeTerms6 `xml:"FXDtls,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`
}

func (s *SwitchOrder4) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchOrder4) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchOrder4) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchOrder4) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchOrder4) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchOrder4) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SwitchOrder4) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SwitchOrder4) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SwitchOrder4) AddExpiryDateTime() *DateAndDateTimeChoice {
	s.ExpiryDateTime = new(DateAndDateTimeChoice)
	return s.ExpiryDateTime
}

func (s *SwitchOrder4) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (s *SwitchOrder4) AddRelatedPartyDetails() *Intermediary8 {
	newValue := new(Intermediary8)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SwitchOrder4) SetCancellationRight(value string) {
	s.CancellationRight = (*CancellationRight1Code)(&value)
}

func (s *SwitchOrder4) SetExtendedCancellationRight(value string) {
	s.ExtendedCancellationRight = (*Extended350Code)(&value)
}

func (s *SwitchOrder4) AddRedemptionLegDetails() *SwitchRedemptionLegOrder3 {
	newValue := new(SwitchRedemptionLegOrder3)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder4) AddSubscriptionLegDetails() *SwitchSubscriptionLegOrder3 {
	newValue := new(SwitchSubscriptionLegOrder3)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchOrder4) AddCashSettlementDetails() *PaymentTransaction25 {
	s.CashSettlementDetails = new(PaymentTransaction25)
	return s.CashSettlementDetails
}

func (s *SwitchOrder4) AddForeignExchangeDetails() *ForeignExchangeTerms6 {
	s.ForeignExchangeDetails = new(ForeignExchangeTerms6)
	return s.ForeignExchangeDetails
}

func (s *SwitchOrder4) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SwitchOrder4) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}
