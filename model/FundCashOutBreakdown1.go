package model

// Breakdown of cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type FundCashOutBreakdown1 struct {

	// Amount of cash flow out, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow out, expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, eg, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Breakdown of the cash movements out of a fund by transaction type, eg, redemption, switch-out.
	InvestmentFundTransactionOutTypeDetails *InvestmentFundTransactionOutType1 `xml:"InvstmtFndTxOutTpDtls,omitempty"`

	// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
	OriginalOrderQuantityDetails *OriginalOrderQuantityType1 `xml:"OrgnlOrdrQtyDtls,omitempty"`

	// Information related to the commission applied to an order, eg, back-end or front-end commission.
	CommissionDetails []*Commission4 `xml:"ComssnDtls,omitempty"`
}

func (f *FundCashOutBreakdown1) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashOutBreakdown1) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashOutBreakdown1) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashOutBreakdown1) AddInvestmentFundTransactionOutTypeDetails() *InvestmentFundTransactionOutType1 {
	f.InvestmentFundTransactionOutTypeDetails = new(InvestmentFundTransactionOutType1)
	return f.InvestmentFundTransactionOutTypeDetails
}

func (f *FundCashOutBreakdown1) AddOriginalOrderQuantityDetails() *OriginalOrderQuantityType1 {
	f.OriginalOrderQuantityDetails = new(OriginalOrderQuantityType1)
	return f.OriginalOrderQuantityDetails
}

func (f *FundCashOutBreakdown1) AddCommissionDetails() *Commission4 {
	newValue := new(Commission4)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}
