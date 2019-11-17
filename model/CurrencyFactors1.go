package model

// Currency specific Factors.
type CurrencyFactors1 struct {

	// Currency of the underlying currency specific amounts and ratios used in the pay-in schedule calculation.
	Currency *CurrencyCode `xml:"Ccy"`

	// Maximum allowed short position in the currency during settlement.
	ShortPositionLimit *ImpliedCurrencyAndAmount `xml:"ShrtPosLmt"`

	// Minimum amount paid in one payment unless the short position is less than the minimum.
	MinimumPayInAmount *ImpliedCurrencyAndAmount `xml:"MinPayInAmt"`

	// Margin used to decrease long positions and increase short positions during the risk calculation.
	VolatilityMargin *PercentageRate `xml:"VoltlyMrgn"`

	// Exchange rate used in the calculation of the pay-in schedule.
	Rate *AgreedRate2 `xml:"Rate,omitempty"`
}

func (c *CurrencyFactors1) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CurrencyFactors1) SetShortPositionLimit(value, currency string) {
	c.ShortPositionLimit = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyFactors1) SetMinimumPayInAmount(value, currency string) {
	c.MinimumPayInAmount = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CurrencyFactors1) SetVolatilityMargin(value string) {
	c.VolatilityMargin = (*PercentageRate)(&value)
}

func (c *CurrencyFactors1) AddRate() *AgreedRate2 {
	c.Rate = new(AgreedRate2)
	return c.Rate
}
