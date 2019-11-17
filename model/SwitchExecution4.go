package model

// Execution of a switch order.
type SwitchExecution4 struct {

	// Unique and unambiguous identifier for a group of individual orders, as assigned by the instructing party. This identifier links the individual orders together.
	MasterReference *Max35Text `xml:"MstrRef,omitempty"`

	// Date and time at which the order was placed by the investor.
	OrderDateTime *ISODateTime `xml:"OrdrDtTm,omitempty"`

	// Account between an investor(s) and a fund manager or a fund. The account can contain holdings in any investment fund or investment fund class managed (or distributed) by the fund manager, within the same fund family.
	InvestmentAccountDetails *InvestmentAccount21 `xml:"InvstmtAcctDtls,omitempty"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous investor's identification of an order. This reference can typically be used in a hub scenario to give the reference of the order as assigned by the underlying client.
	ClientReference *Max35Text `xml:"ClntRef,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be redeemed.
	TotalRedemptionAmount *ActiveCurrencyAndAmount `xml:"TtlRedAmt,omitempty"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	TotalSubscriptionAmount *ActiveCurrencyAndAmount `xml:"TtlSbcptAmt,omitempty"`

	// Confirmation of the information about parties related to the transaction.
	RelatedPartyDetails []*Intermediary9 `xml:"RltdPtyDtls,omitempty"`

	// Future date at which the investor requests the order to be executed.
	// The specification of a requested future trade date is not allowed in some markets. The date must be a date in the future.
	RequestedFutureTradeDate *ISODate `xml:"ReqdFutrTradDt,omitempty"`

	// Additional amount of money paid by the investor in addition to the switch redemption amount.
	AdditionalCashIn *ActiveCurrencyAndAmount `xml:"AddtlCshIn,omitempty"`

	// Amount of money that results from a switch-out, that is not reinvested in another investment fund, and is repaid to the investor.
	ResultingCashOut *ActiveCurrencyAndAmount `xml:"RsltgCshOut,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Date on which cash is available.
	CashSettlementDate *ISODate `xml:"CshSttlmDt,omitempty"`

	// Method by which the transaction is settled.
	SettlementMethod *DeliveryReceiptType2Code `xml:"SttlmMtd,omitempty"`

	// Specifies that the execution was subject to best execution rules as defined by MiFID.
	BestExecution *BestExecution1Code `xml:"BestExctn,omitempty"`

	// Redemption leg of a switch order execution.
	RedemptionLegDetails []*SwitchRedemptionLegExecution3 `xml:"RedLegDtls"`

	// Subscription leg of a switch order execution.
	SubscriptionLegDetails []*SwitchSubscriptionLegExecution3 `xml:"SbcptLegDtls"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction26 `xml:"CshSttlmDtls,omitempty"`

	// Currency exchange related to the execution of an investment fund order.
	ForeignExchangeDetails []*ForeignExchangeTerms7 `xml:"FXDtls,omitempty"`

	// Specifies if advice has been received from an independent financial advisor.
	FinancialAdvice *FinancialAdvice1Code `xml:"FinAdvc,omitempty"`

	// Specifies whether the trade is negotiated.
	NegotiatedTrade *NegotiatedTrade1Code `xml:"NgtdTrad,omitempty"`

	// Specifies whether the order execution confirmation is late.
	LateReport *LateReport1Code `xml:"LateRpt,omitempty"`
}

func (s *SwitchExecution4) SetMasterReference(value string) {
	s.MasterReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetOrderDateTime(value string) {
	s.OrderDateTime = (*ISODateTime)(&value)
}

func (s *SwitchExecution4) AddInvestmentAccountDetails() *InvestmentAccount21 {
	s.InvestmentAccountDetails = new(InvestmentAccount21)
	return s.InvestmentAccountDetails
}

func (s *SwitchExecution4) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetClientReference(value string) {
	s.ClientReference = (*Max35Text)(&value)
}

func (s *SwitchExecution4) SetTotalRedemptionAmount(value, currency string) {
	s.TotalRedemptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetTotalSubscriptionAmount(value, currency string) {
	s.TotalSubscriptionAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) AddRelatedPartyDetails() *Intermediary9 {
	newValue := new(Intermediary9)
	s.RelatedPartyDetails = append(s.RelatedPartyDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) SetRequestedFutureTradeDate(value string) {
	s.RequestedFutureTradeDate = (*ISODate)(&value)
}

func (s *SwitchExecution4) SetAdditionalCashIn(value, currency string) {
	s.AdditionalCashIn = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetResultingCashOut(value, currency string) {
	s.ResultingCashOut = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetSettlementAmount(value, currency string) {
	s.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchExecution4) SetCashSettlementDate(value string) {
	s.CashSettlementDate = (*ISODate)(&value)
}

func (s *SwitchExecution4) SetSettlementMethod(value string) {
	s.SettlementMethod = (*DeliveryReceiptType2Code)(&value)
}

func (s *SwitchExecution4) SetBestExecution(value string) {
	s.BestExecution = (*BestExecution1Code)(&value)
}

func (s *SwitchExecution4) AddRedemptionLegDetails() *SwitchRedemptionLegExecution3 {
	newValue := new(SwitchRedemptionLegExecution3)
	s.RedemptionLegDetails = append(s.RedemptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) AddSubscriptionLegDetails() *SwitchSubscriptionLegExecution3 {
	newValue := new(SwitchSubscriptionLegExecution3)
	s.SubscriptionLegDetails = append(s.SubscriptionLegDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) AddCashSettlementDetails() *PaymentTransaction26 {
	s.CashSettlementDetails = new(PaymentTransaction26)
	return s.CashSettlementDetails
}

func (s *SwitchExecution4) AddForeignExchangeDetails() *ForeignExchangeTerms7 {
	newValue := new(ForeignExchangeTerms7)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SwitchExecution4) SetFinancialAdvice(value string) {
	s.FinancialAdvice = (*FinancialAdvice1Code)(&value)
}

func (s *SwitchExecution4) SetNegotiatedTrade(value string) {
	s.NegotiatedTrade = (*NegotiatedTrade1Code)(&value)
}

func (s *SwitchExecution4) SetLateReport(value string) {
	s.LateReport = (*LateReport1Code)(&value)
}
