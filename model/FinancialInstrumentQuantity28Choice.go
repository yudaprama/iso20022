package model

// Choice between formats for the specification of the number of units, amount of money or percentage..
type FinancialInstrumentQuantity28Choice struct {

	// Number of investment fund units redeemed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Amount of money to be redeemed from the fund.
	// Gross Amount = Quantity * Price.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money to be received following redemption of fund units.
	// Net Amount = (Quantity * Price) - (Fees + Taxes).
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Portion of the investor's holdings to be redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`
}

func (f *FinancialInstrumentQuantity28Choice) SetUnitsNumber(value string) {
	f.UnitsNumber = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity28Choice) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity28Choice) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity28Choice) SetHoldingsRedemptionRate(value string) {
	f.HoldingsRedemptionRate = (*PercentageRate)(&value)
}
