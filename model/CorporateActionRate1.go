package model

// Specifies rates.
type CorporateActionRate1 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat1Choice `xml:"Intrst,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat1Choice `xml:"RltdIndx,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	PercentageSought *RateFormat1Choice `xml:"PctgSght,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountToMarket *RateFormat1Choice `xml:"RinvstmtDscntToMkt,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat1Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *AmountAndRateFormat3Choice `xml:"BidIntrvl,omitempty"`

	// Rate used to calculate the amount of the charges/fees that cannot be categorised.
	Charges *RateAndAmountFormat1Choice `xml:"Chrgs,omitempty"`
}

func (c *CorporateActionRate1) AddInterest() *RateAndAmountFormat1Choice {
	c.Interest = new(RateAndAmountFormat1Choice)
	return c.Interest
}

func (c *CorporateActionRate1) AddRelatedIndex() *RateFormat1Choice {
	c.RelatedIndex = new(RateFormat1Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate1) AddPercentageSought() *RateFormat1Choice {
	c.PercentageSought = new(RateFormat1Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate1) AddReinvestmentDiscountToMarket() *RateFormat1Choice {
	c.ReinvestmentDiscountToMarket = new(RateFormat1Choice)
	return c.ReinvestmentDiscountToMarket
}

func (c *CorporateActionRate1) AddSpread() *RateFormat1Choice {
	c.Spread = new(RateFormat1Choice)
	return c.Spread
}

func (c *CorporateActionRate1) AddBidInterval() *AmountAndRateFormat3Choice {
	c.BidInterval = new(AmountAndRateFormat3Choice)
	return c.BidInterval
}

func (c *CorporateActionRate1) AddCharges() *RateAndAmountFormat1Choice {
	c.Charges = new(RateAndAmountFormat1Choice)
	return c.Charges
}
