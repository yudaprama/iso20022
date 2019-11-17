package model

// Extract of trade data for an investment fund order.
type FundOrderData1 struct {

	// Account information of the individual order instruction for which the status is given.
	InvestmentAccountDetails *InvestmentAccount13 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order instruction for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument10 `xml:"FinInstrmDtls,omitempty"`

	// Quantity of investment fund units subscribed or redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be sold or subscribed to.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// Amount of money used to derive the quantity of investment fund units to be sold or subscribed to, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Currency in which the rate of exchange is expressed in a currency exchange. In the example 1GBP = xxxCUR, the unit currency is GBP.
	UnitCurrency *ActiveOrHistoricCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the base currency is converted, in a currency exchange.
	QuotedCurrency *ActiveOrHistoricCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData1) AddInvestmentAccountDetails() *InvestmentAccount13 {
	f.InvestmentAccountDetails = new(InvestmentAccount13)
	return f.InvestmentAccountDetails
}

func (f *FundOrderData1) AddFinancialInstrumentDetails() *FinancialInstrument10 {
	f.FinancialInstrumentDetails = new(FinancialInstrument10)
	return f.FinancialInstrumentDetails
}

func (f *FundOrderData1) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FundOrderData1) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData1) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData1) SetHoldingsRedemptionRate(value string) {
	f.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (f *FundOrderData1) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData1) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}

func (f *FundOrderData1) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveOrHistoricCurrencyCode)(&value)
}
