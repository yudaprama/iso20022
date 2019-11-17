package model

// Breakdown of cash movements out of a fund as a result of investment funds transactions, eg, redemptions or switch-out.
type FundCashOutBreakdown2 struct {

	// Amount of cash flow out, expressed as an amount of money.
	Amount *ActiveOrHistoricCurrencyAndAmount `xml:"Amt,omitempty"`

	// Amount of the cash flow out, expressed as a number of units.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Indicates whether the cash flow is an item that did not appear on the previously sent report, eg, because it was received close to cut-off time.
	NewAmountIndicator *YesNoIndicator `xml:"NewAmtInd,omitempty"`

	// Breakdown of the cash movements out of a fund by transaction type, eg, redemption, switch-out.
	InvestmentFundTransactionOutType *InvestmentFundTransactionOutType1Code `xml:"InvstmtFndTxOutTp"`

	// Breakdown of the cash movements out of a fund by transaction type, eg, redemption, switch-out.
	ExtendedInvestmentFundTransactionOutType *Extended350Code `xml:"XtndedInvstmtFndTxOutTp"`

	// Breakdown of the cash movements into a fund by order type, eg, order by quantity of units or amount of money.
	OriginalOrderQuantityType *OrderQuantityType2Code `xml:"OrgnlOrdrQtyTp"`

	// Breakdown of the cash movements into a fund by transaction type, eg, subscription, switch-in.
	ExtendedOriginalOrderQuantityType *Extended350Code `xml:"XtndedOrgnlOrdrQtyTp"`

	// Charge for the placement of an order.
	ChargeDetails []*Charge16 `xml:"ChrgDtls,omitempty"`

	// Information related to the commission applied to an order, eg, back-end or front-end commission.
	CommissionDetails []*Commission9 `xml:"ComssnDtls,omitempty"`
}

func (f *FundCashOutBreakdown2) SetAmount(value, currency string) {
	f.Amount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundCashOutBreakdown2) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundCashOutBreakdown2) SetNewAmountIndicator(value string) {
	f.NewAmountIndicator = (*YesNoIndicator)(&value)
}

func (f *FundCashOutBreakdown2) SetInvestmentFundTransactionOutType(value string) {
	f.InvestmentFundTransactionOutType = (*InvestmentFundTransactionOutType1Code)(&value)
}

func (f *FundCashOutBreakdown2) SetExtendedInvestmentFundTransactionOutType(value string) {
	f.ExtendedInvestmentFundTransactionOutType = (*Extended350Code)(&value)
}

func (f *FundCashOutBreakdown2) SetOriginalOrderQuantityType(value string) {
	f.OriginalOrderQuantityType = (*OrderQuantityType2Code)(&value)
}

func (f *FundCashOutBreakdown2) SetExtendedOriginalOrderQuantityType(value string) {
	f.ExtendedOriginalOrderQuantityType = (*Extended350Code)(&value)
}

func (f *FundCashOutBreakdown2) AddChargeDetails() *Charge16 {
	newValue := new(Charge16)
	f.ChargeDetails = append(f.ChargeDetails, newValue)
	return newValue
}

func (f *FundCashOutBreakdown2) AddCommissionDetails() *Commission9 {
	newValue := new(Commission9)
	f.CommissionDetails = append(f.CommissionDetails, newValue)
	return newValue
}
