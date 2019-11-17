package model

// Set of elements qualifying the interest rate.
type Rate1 struct {

	// Percentage charged for the use of an amount of money, usually expressed at an annual rate. The interest rate is the ratio of the amount of interest paid during a certain period of time compared to the principal amount of the interest bearing financial instrument.
	// Example percentage rate : Rate expressed as a percentage, ie, in hundredths, eg, 0.7 is 7/10 of a percent, and 7.0 is 7%.
	// Example Textual rate : Rate is expressed as a text.
	Rate *RateTypeChoice `xml:"Rate"`

	// An amount range where the interest rate is applicable
	ValidityRange *CurrencyAndAmountRange `xml:"VldtyRg,omitempty"`
}

func (r *Rate1) AddRate() *RateTypeChoice {
	r.Rate = new(RateTypeChoice)
	return r.Rate
}

func (r *Rate1) AddValidityRange() *CurrencyAndAmountRange {
	r.ValidityRange = new(CurrencyAndAmountRange)
	return r.ValidityRange
}
