package model

// Execution of the subscription part, in a switch between investment funds or investment fund classes.
type SwitchSubscriptionLegExecution2 struct {

	// Unique technical identifier for an instance of a leg within a switch.
	LegIdentification *Max35Text `xml:"LegId,omitempty"`

	// Unique identifier for an instance of a leg execution within a switch confirmation.
	LegExecutionIdentification *Max35Text `xml:"LegExctnId,omitempty"`

	// Security that is a sub-set of an investment fund, and is governed by the same investment fund policy, eg, dividend option or valuation currency.
	FinancialInstrumentDetails *FinancialInstrument6 `xml:"FinInstrmDtls"`

	// Number of investment fund units subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Net amount of money invested in a specific financial instrument by an investor, expressed in the currency requested by the investor.
	NetAmount *ActiveCurrencyAndAmount `xml:"NetAmt"`

	// Gross amount of money invested in a specific financial instrument by an investor, expressed in the currency requested by the investor.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Date and time at which a price is applied, according to the terms stated in the prospectus.
	TradeDateTime *DateAndDateTimeChoice `xml:"TradDtTm"`

	// Price at which the order was executed.
	PriceDetails *UnitPrice5 `xml:"PricDtls"`

	// Indicates whether the dividend is included, ie, cum-dividend, in the executed price. When the dividend is not included, the price will be ex-dividend.
	CumDividendIndicator *YesNoIndicator `xml:"CumDvddInd"`

	// Part of the price deemed as accrued income or profit rather than capital. The interim profit amount is used for tax purposes.
	InterimProfitAmount *ProfitAndLoss1Choice `xml:"IntrmPrftAmt,omitempty"`

	// Dividend option chosen by the account owner based on the options offered in the prospectus.
	IncomePreference *IncomePreference1Code `xml:"IncmPref,omitempty"`

	// Currency requested for settlement of cash proceeds.
	RequestedSettlementCurrency *CurrencyCode `xml:"ReqdSttlmCcy,omitempty"`

	// Currency to be used for pricing the fund. This currency must be among the set of currencies in which the price may be expressed, as stated in the prospectus.
	RequestedNAVCurrency *CurrencyCode `xml:"ReqdNAVCcy,omitempty"`

	// Charge for the execution of an order.
	ChargeGeneralDetails *TotalCharges2 `xml:"ChrgGnlDtls,omitempty"`

	// Commission for the execution of an investment fund order.
	CommissionGeneralDetails *TotalCommissions2 `xml:"ComssnGnlDtls,omitempty"`

	// Tax applicable to execution of an investment fund order.
	TaxGeneralDetails *TotalTaxes2 `xml:"TaxGnlDtls,omitempty"`

	// Parameters used to execute the settlement of an investment fund order.
	SettlementAndCustodyDetails *FundSettlementParameters3 `xml:"SttlmAndCtdyDtls,omitempty"`

	// Indicates whether the financial instrument is to be physically delivered.
	PhysicalDeliveryIndicator *YesNoIndicator `xml:"PhysDlvryInd"`

	// Information related to physical delivery of the securities.
	PhysicalDeliveryDetails *DeliveryParameters3 `xml:"PhysDlvryDtls,omitempty"`
}

func (s *SwitchSubscriptionLegExecution2) SetLegIdentification(value string) {
	s.LegIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution2) SetLegExecutionIdentification(value string) {
	s.LegExecutionIdentification = (*Max35Text)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddFinancialInstrumentDetails() *FinancialInstrument6 {
	s.FinancialInstrumentDetails = new(FinancialInstrument6)
	return s.FinancialInstrumentDetails
}

func (s *SwitchSubscriptionLegExecution2) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	s.UnitsNumber = new(FinancialInstrumentQuantity1)
	return s.UnitsNumber
}

func (s *SwitchSubscriptionLegExecution2) SetNetAmount(value, currency string) {
	s.NetAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution2) SetGrossAmount(value, currency string) {
	s.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (s *SwitchSubscriptionLegExecution2) AddTradeDateTime() *DateAndDateTimeChoice {
	s.TradeDateTime = new(DateAndDateTimeChoice)
	return s.TradeDateTime
}

func (s *SwitchSubscriptionLegExecution2) AddPriceDetails() *UnitPrice5 {
	s.PriceDetails = new(UnitPrice5)
	return s.PriceDetails
}

func (s *SwitchSubscriptionLegExecution2) SetCumDividendIndicator(value string) {
	s.CumDividendIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddInterimProfitAmount() *ProfitAndLoss1Choice {
	s.InterimProfitAmount = new(ProfitAndLoss1Choice)
	return s.InterimProfitAmount
}

func (s *SwitchSubscriptionLegExecution2) SetIncomePreference(value string) {
	s.IncomePreference = (*IncomePreference1Code)(&value)
}

func (s *SwitchSubscriptionLegExecution2) SetRequestedSettlementCurrency(value string) {
	s.RequestedSettlementCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution2) SetRequestedNAVCurrency(value string) {
	s.RequestedNAVCurrency = (*CurrencyCode)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddChargeGeneralDetails() *TotalCharges2 {
	s.ChargeGeneralDetails = new(TotalCharges2)
	return s.ChargeGeneralDetails
}

func (s *SwitchSubscriptionLegExecution2) AddCommissionGeneralDetails() *TotalCommissions2 {
	s.CommissionGeneralDetails = new(TotalCommissions2)
	return s.CommissionGeneralDetails
}

func (s *SwitchSubscriptionLegExecution2) AddTaxGeneralDetails() *TotalTaxes2 {
	s.TaxGeneralDetails = new(TotalTaxes2)
	return s.TaxGeneralDetails
}

func (s *SwitchSubscriptionLegExecution2) AddSettlementAndCustodyDetails() *FundSettlementParameters3 {
	s.SettlementAndCustodyDetails = new(FundSettlementParameters3)
	return s.SettlementAndCustodyDetails
}

func (s *SwitchSubscriptionLegExecution2) SetPhysicalDeliveryIndicator(value string) {
	s.PhysicalDeliveryIndicator = (*YesNoIndicator)(&value)
}

func (s *SwitchSubscriptionLegExecution2) AddPhysicalDeliveryDetails() *DeliveryParameters3 {
	s.PhysicalDeliveryDetails = new(DeliveryParameters3)
	return s.PhysicalDeliveryDetails
}
