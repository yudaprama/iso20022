package model

// Choice between formats for the specification of the number of units, amount of money or percentage.
type FinancialInstrumentQuantity26Choice struct {

	// Number of investment fund units to be subscribed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Percentage of the total redemption amount used for the subscription in an investment fund or investment fund class.
	PercentageOfTotalRedemptionAmount *PercentageRate `xml:"PctgOfTtlRedAmt"`

	// Amount of money to be invested in the fund.
	// Net Amount = Quantity * Price.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`
}

func (f *FinancialInstrumentQuantity26Choice) SetUnitsNumber(value string) {
	f.UnitsNumber = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity26Choice) SetPercentageOfTotalRedemptionAmount(value string) {
	f.PercentageOfTotalRedemptionAmount = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentQuantity26Choice) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity26Choice) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
