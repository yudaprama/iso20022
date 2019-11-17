package model

// Execution of a subscription order.
type SubscriptionExecution4 struct {

	// Unique and unambiguous identifier for an order, as assigned by the instructing party.
	OrderReference *Max35Text `xml:"OrdrRef"`

	// Unique and unambiguous identifier for an order execution, as assigned by a confirming party.
	DealReference *Max35Text `xml:"DealRef"`

	// Specifies the category of the investment fund order.
	OrderType []*FundOrderType1 `xml:"OrdrTp,omitempty"`

	// Investment fund class to which an investment fund order execution is related.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Indicates the rounding direction applied to nearest unit.
	Rounding *RoundingDirection2Code `xml:"Rndg,omitempty"`

	// Net amount of money invested in a specific financial instrument by an investor, expressed in the currency requested by the investor.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Amount of money invested in a specific financial instrument by an investor, including all charges, commissions, and tax, expressed in the currency requested by the investor.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice5 `xml:"PricDtls"`

	// Indicates whether the order has been partially executed, ie, the confirmed quantity does not match the ordered quantity for a given financial instrument.
	PartiallyExecutedIndicator *YesNoIndicator `xml:"PrtlyExctdInd"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Information needed to process a currency exchange or conversion.
	ForeignExchangeDetails []*ForeignExchangeTerms4 `xml:"FXDtls,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Reference of a letter of intent program, in which sales commissions are reduced based on the aggregate of a customer's actual purchase and anticipated purchases, over a specific period of time, and as agreed by the customer. A letter of intent program is mainly used in the US market.
	LetterIntentReference *Max35Text `xml:"LttrInttRef,omitempty"`

	// Reference of an accumulation right program, in which sales commissions are based on a customer's present purchases of shares and the aggregate quantity previously purchased by the customer. An accumulation rights program is mainly used in the US market.
	AccumulationRightReference *Max35Text `xml:"AcmltnRghtRef,omitempty"`

	// Amount of money associated with a service.
	ChargeGeneralDetails *TotalCharges2 `xml:"ChrgGnlDtls,omitempty"`

	// Amount of money due to a party as compensation for a service.
	CommissionGeneralDetails *TotalCommissions2 `xml:"ComssnGnlDtls,omitempty"`

	// Tax related to an investment fund order.
	TaxGeneralDetails *TotalTaxes2 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Return of cash that has been overpaid for a subscription.
	Refund *ActiveCurrencyAndAmount `xml:"Rfnd,omitempty"`

	// Interest received when a subscription amount is paid in advance and then invested by the fund.
	SubscriptionInterest *ActiveCurrencyAndAmount `xml:"SbcptIntrst,omitempty"`

	// Payment transaction resulting from the investment fund order execution.
	CashSettlementDetails *PaymentTransaction13 `xml:"CshSttlmDtls,omitempty"`
}

func (s *SubscriptionExecution4) SetOrderReference(value string) {
	s.OrderReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) SetDealReference(value string) {
	s.DealReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) AddOrderType() *FundOrderType1 {
	newValue := new(FundOrderType1)
	s.OrderType = append(s.OrderType, newValue)
	return newValue
}

func (s *SubscriptionExecution4) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SubscriptionExecution4) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SubscriptionExecution4) SetRounding(value string) {
	s.Rounding = (*RoundingDirection2Code)(&value)
}

func (s *SubscriptionExecution4) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SubscriptionExecution4) AddPriceDetails() *UnitPrice5 {
	s.PriceDetails = new(UnitPrice5)
	return s.PriceDetails
}

func (s *SubscriptionExecution4) SetPartiallyExecutedIndicator(value string) {
	s.PartiallyExecutedIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution4) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution4) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SubscriptionExecution4) AddForeignExchangeDetails() *ForeignExchangeTerms4 {
	newValue := new(ForeignExchangeTerms4)
	s.ForeignExchangeDetails = append(s.ForeignExchangeDetails, newValue)
	return newValue
}

func (s *SubscriptionExecution4) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SubscriptionExecution4) SetLetterIntentReference(value string) {
	s.LetterIntentReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) SetAccumulationRightReference(value string) {
	s.AccumulationRightReference = (*Max35Text)(&value)
}

func (s *SubscriptionExecution4) AddChargeGeneralDetails() *TotalCharges2 {
	s.ChargeGeneralDetails = new(TotalCharges2)
	return s.ChargeGeneralDetails
}

func (s *SubscriptionExecution4) AddCommissionGeneralDetails() *TotalCommissions2 {
	s.CommissionGeneralDetails = new(TotalCommissions2)
	return s.CommissionGeneralDetails
}

func (s *SubscriptionExecution4) AddTaxGeneralDetails() *TotalTaxes2 {
	s.TaxGeneralDetails = new(TotalTaxes2)
	return s.TaxGeneralDetails
}

func (s *SubscriptionExecution4) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SubscriptionExecution4) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SubscriptionExecution4) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}

func (s *SubscriptionExecution4) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionExecution4) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SubscriptionExecution4) SetRefund(value, currency string) {
	s.Refund = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) SetSubscriptionInterest(value, currency string) {
	s.SubscriptionInterest = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SubscriptionExecution4) AddCashSettlementDetails() *PaymentTransaction13 {
	s.CashSettlementDetails = new(PaymentTransaction13)
	return s.CashSettlementDetails
}
