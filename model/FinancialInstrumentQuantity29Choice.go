package model

// Choice between formats for the specification of the number of units, amount of money or percentage.
type FinancialInstrumentQuantity29Choice struct {

	// Number of investment fund units redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Percentage of the total switch amount (buy-driven) to be invested in a particular investment fund or investment fund class.
	PercentageOfTotalSubscriptionAmount *PercentageRate `xml:"PctgOfTtlSbcptAmt"`

	// Amount of money to be received following redemption of fund units.
	// Net Amount = (Quantity * Price) - (Fees + Taxes).
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Amount of money to be redeemed from the fund.
	// Gross Amount = Quantity * Price.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Portion of the investor's holdings to be redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`
}

func (f *FinancialInstrumentQuantity29Choice) SetUnitsNumber(value string) {
	f.UnitsNumber = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity29Choice) SetPercentageOfTotalSubscriptionAmount(value string) {
	f.PercentageOfTotalSubscriptionAmount = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentQuantity29Choice) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity29Choice) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity29Choice) SetHoldingsRedemptionRate(value string) {
	f.HoldingsRedemptionRate = (*PercentageRate)(&value)
}
