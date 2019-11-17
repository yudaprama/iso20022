package model

// Choice between ways to express the quantity of the financial instrument to be subscribed.
type FinancialInstrumentQuantity6Choice struct {

	// Quantity of investment fund units to be subscribed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Percentage of the total redemption amount used for the subscription in an investment fund or investment fund class.
	PercentageOfTotalRedemptionAmount *PercentageRate `xml:"PctgOfTtlRedAmt"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Amount of money used to determine the quantity of investment fund units to be subscribed, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`
}

func (f *FinancialInstrumentQuantity6Choice) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FinancialInstrumentQuantity6Choice) SetPercentageOfTotalRedemptionAmount(value string) {
	f.PercentageOfTotalRedemptionAmount = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentQuantity6Choice) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity6Choice) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}
