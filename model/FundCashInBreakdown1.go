package model

// Breakdown of cash movements into a fund as a result of investment funds transactions, eg, subscriptions or switch-in.
type FundCashInBreakdown1 struct {

	// Amount of cash flow in, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow in,  expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, eg, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Breakdown of the cash movements into a fund by transaction type, eg, subscription, switch-in.
	InvestmentFundTransactionInTypeDetails *InvestmentFundTransactionInType1 `xml:"InvstmtFndTxInTpDtls,omitempty"`

	// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
	OriginalOrderQuantityDetails *OriginalOrderQuantityType1 `xml:"OrgnlOrdrQtyDtls,omitempty"`

	// Information related to the commission applied to an order, eg, back-end or front-end commission.
	CommissionDetails []*Commission4 `xml:"ComssnDtls,omitempty"`
}

func (f *FundCashInBreakdown1) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashInBreakdown1) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashInBreakdown1) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashInBreakdown1) AddInvestmentFundTransactionInTypeDetails() *InvestmentFundTransactionInType1 {
	f.InvestmentFundTransactionInTypeDetails = new(InvestmentFundTransactionInType1)
	return f.InvestmentFundTransactionInTypeDetails
}

func (f *FundCashInBreakdown1) AddOriginalOrderQuantityDetails() *OriginalOrderQuantityType1 {
	f.OriginalOrderQuantityDetails = new(OriginalOrderQuantityType1)
	return f.OriginalOrderQuantityDetails
}

func (f *FundCashInBreakdown1) AddCommissionDetails() *Commission4 {
	newValue := new(Commission4)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}
