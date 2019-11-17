package model

// Breakdown of cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type FundCashInBreakdown3 struct {

	// Amount of cash flow in, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow in,  expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, for example, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Type of transaction that resulted in the cash-in movement, for example, subscription, switch-in.
	InvestmentFundTransactionInType *InvestmentFundTransactionInType1Choice `xml:"InvstmtFndTxInTp"`

	// Specifies how the original order was expressed that resulted in the cash-in movement, that is cash or units.
	OriginalOrderQuantityType *QuantityType1Choice `xml:"OrgnlOrdrQtyTp"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge26 `xml:"ChrgDtls,omitempty"`

	// Information related to the commission applied to an order, for example, back-end or front-end commission.
	CommissionDetails []*Commission21 `xml:"ComssnDtls,omitempty"`

	// Settlement currency for the transaction.
	SettlementCurrency *ActiveCurrencyCode `xml:"SttlmCcy,omitempty"`
}

func (f *FundCashInBreakdown3) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashInBreakdown3) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashInBreakdown3) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashInBreakdown3) AddInvestmentFundTransactionInType() *InvestmentFundTransactionInType1Choice {
	f.InvestmentFundTransactionInType = new(InvestmentFundTransactionInType1Choice)
	return f.InvestmentFundTransactionInType
}

func (f *FundCashInBreakdown3) AddOriginalOrderQuantityType() *QuantityType1Choice {
	f.OriginalOrderQuantityType = new(QuantityType1Choice)
	return f.OriginalOrderQuantityType
}

func (f *FundCashInBreakdown3) AddChargeDetails() *Charge26 {
	newValue := new(Charge26)
	f.ChargeDetails = append(f.ChargeDetails, newValue)
	return newValue
}

func (f *FundCashInBreakdown3) AddCommissionDetails() *Commission21 {
	newValue := new(Commission21)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}

func (f *FundCashInBreakdown3) SetSettlementCurrency(value string) {
	f.SettlementCurrency = (*ActiveCurrencyCode)(&value)
}
