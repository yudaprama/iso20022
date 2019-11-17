package model

// Choice between ways to express the quantity of the financial instrument to be redeemed.
type FinancialInstrumentQuantity5Choice struct {

	// Quantity of investment fund units redeemed.
	UnitsNumber *FinancialInstrumentQuantity1 `xml:"UnitsNb"`

	// Percentage of the total switch amount (buy-driven) to be invested in a particular investment fund or investment fund class.
	PercentageOfTotalSubscriptionAmount *PercentageRate `xml:"PctgOfTtlSbcptAmt"`

	// Amount of money used to derive the quantity of investment fund units to be sold.
	NetAmount *ActiveOrHistoricCurrencyAndAmount `xml:"NetAmt"`

	// Amount of money used to derive the quantity of investment fund units to be sold, including all charges, commissions, and tax.
	GrossAmount *ActiveOrHistoricCurrencyAndAmount `xml:"GrssAmt"`

	// Portion of the investor's holdings, in a specific investment fund/ fund class, that is redeemed.
	HoldingsRedemptionRate *PercentageRate `xml:"HldgsRedRate"`
}

func (f *FinancialInstrumentQuantity5Choice) AddUnitsNumber() *FinancialInstrumentQuantity1 {
	f.UnitsNumber = new(FinancialInstrumentQuantity1)
	return f.UnitsNumber
}

func (f *FinancialInstrumentQuantity5Choice) SetPercentageOfTotalSubscriptionAmount(value string) {
	f.PercentageOfTotalSubscriptionAmount = (*PercentageRate)(&value)
}

func (f *FinancialInstrumentQuantity5Choice) SetNetAmount(value, currency string) {
	f.NetAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity5Choice) SetGrossAmount(value, currency string) {
	f.GrossAmount = NewActiveOrHistoricCurrencyAndAmount(value, currency)
}

func (f *FinancialInstrumentQuantity5Choice) SetHoldingsRedemptionRate(value string) {
	f.HoldingsRedemptionRate = (*PercentageRate)(&value)
}
