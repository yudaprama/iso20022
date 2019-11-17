package model

// Extract of trade data for an investment fund order.
type FundOrderData5 struct {

	// Account information of the individual order instruction for which the status is given.
	InvestmentAccountDetails *InvestmentAccount58 `xml:"InvstmtAcctDtls,omitempty"`

	// Financial instrument information of the individual order instruction for which the status is given.
	FinancialInstrumentDetails *FinancialInstrument57 `xml:"FinInstrmDtls,omitempty"`

	// Number of investment fund units subscribed or redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb,omitempty"`

	// When the status message is used for a subscription, this is the amount of money to be invested in the fund.
	// Net Amount = Quantity * Price.
	// When the status message is used for a redemption, this is the amount of money to be received following redemption of fund units.
	// Net Amount = (Quantity * Price) - (Fees + Taxes).
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt,omitempty"`

	// When the status message is used for a subscription, this is the amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	// When the status message is used for a redemption, this is the amount of money to be redeemed from the fund.
	// Gross Amount = Quantity * Price.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt,omitempty"`

	// Portion of the investor's holdings, in a specific investment fund/fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate,omitempty"`

	// Total amount of money paid /to be paid or received in exchange for the financial instrument in the individual order.
	SettlementAmount *ActiveCurrencyAndAmount `xml:"SttlmAmt,omitempty"`

	// Currency from which the quoted currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	UnitCurrency *ActiveCurrencyCode `xml:"UnitCcy,omitempty"`

	// Currency into which the unit currency is converted in an exchange rate calculation.
	// 1 x <UnitCcy> = <XchgRate> x <QtdCcy>.
	QuotedCurrency *ActiveCurrencyCode `xml:"QtdCcy,omitempty"`
}

func (f *FundOrderData5) AddInvestmentAccountDetails() *InvestmentAccount58 {
	f.InvestmentAccountDetails = new(InvestmentAccount58)
	return f.InvestmentAccountDetails
}

func (f *FundOrderData5) AddFinancialInstrumentDetails() *FinancialInstrument57 {
	f.FinancialInstrumentDetails = new(FinancialInstrument57)
	return f.FinancialInstrumentDetails
}

func (f *FundOrderData5) SetUnitsNumber(value string) {
	f.UnitsNumber = (*DecimalNumber)(&value)
}

func (f *FundOrderData5) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData5) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FundOrderData5) SetHoldingsRedemptionRate(value string) {
	f.HoldingsRedemptionRate = (*PercentageRate)(&value)
}

func (f *FundOrderData5) SetSettlementAmount(value, currency string) {
	f.SettlementAmount = NewActiveCurrencyAndAmount(value, currency)
}

func (f *FundOrderData5) SetUnitCurrency(value string) {
	f.UnitCurrency = (*ActiveCurrencyCode)(&value)
}

func (f *FundOrderData5) SetQuotedCurrency(value string) {
	f.QuotedCurrency = (*ActiveCurrencyCode)(&value)
}
