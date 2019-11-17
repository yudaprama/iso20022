package model

// Choice between ways to express the quantity of the financial instrument to be subscribed.
type FinancialInstrumentQuantity4Choice struct {

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Percentage of the total redemption amount used for the subscription in an investment fund or investment fund class.
	PercentageOfTotalRedemptionAmount *PercentageRate `xml:"PctgOfTtlRedAmt"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	NetAmount *CurrencyAndAmount `xml:"NetAmt"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed, including all charges, commissions, and tax.
	GrossAmount *ActiveCurrencyAndAmount `xml:"GrssAmt"`
}

func (f *FinancialInstrumentQuantity4Choice) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FinancialInstrumentQuantity4Choice) SetPercentageOfTotalRedemptionAmount(value string) {
	f.PercentageOfTotalRedemptionAmount = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentQuantity4Choice) SetNetAmount(value, currency string) {
	f.NetAmount = NewCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity4Choice) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveCurrencyAndAmount(value, currency)
}
