package model

// Choice between formats for the specification of the number of units or amount of money.
type FinancialInstrumentQuantity27Choice struct {

	// Number of investment fund units to be subscribed.
	UnitsNumber *DecimalNumber `xml:"UnitsNb"`

	// Amount of money to be paid by the investor when subscribing to fund units.
	// Gross amount = (Quantity * Price) + (Fees + Taxes).
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Amount of money to be invested in the fund.
	// Net Amount = Quantity * Price.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`
}

func (f *FinancialInstrumentQuantity27Choice) SetUnitsNumber(value string) {
	f.UnitsNumber = (*DecimalNumber)(&value)
}

func (f *FinancialInstrumentQuantity27Choice) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity27Choice) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
