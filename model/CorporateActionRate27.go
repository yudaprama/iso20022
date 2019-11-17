package model

// Specifies rates of a corporate action.
type CorporateActionRate27 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat14Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event.
	PercentageSought *RateFormat7Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat6Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat6Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat14Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat6Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the interest payment and the expected or scheduled rate of the interest payment .
	InterestShortfall *RateAndAmountFormat12Choice `xml:"IntrstShrtfll,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the capital or principal repayment and the scheduled capital repayment.
	RealisedLoss *RateAndAmountFormat12Choice `xml:"RealsdLoss,omitempty"`
}

func (c *CorporateActionRate27) AddInterest() *RateAndAmountFormat14Choice {
	c.Interest = new(RateAndAmountFormat14Choice)
	return c.Interest
}

func (c *CorporateActionRate27) AddPercentageSought() *RateFormat7Choice {
	c.PercentageSought = new(RateFormat7Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate27) AddRelatedIndex() *RateFormat6Choice {
	c.RelatedIndex = new(RateFormat6Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate27) AddSpread() *RateFormat6Choice {
	c.Spread = new(RateFormat6Choice)
	return c.Spread
}

func (c *CorporateActionRate27) AddBidInterval() *RateAndAmountFormat14Choice {
	c.BidInterval = new(RateAndAmountFormat14Choice)
	return c.BidInterval
}

func (c *CorporateActionRate27) AddPreviousFactor() *RateFormat3Choice {
	c.PreviousFactor = new(RateFormat3Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate27) AddNextFactor() *RateFormat3Choice {
	c.NextFactor = new(RateFormat3Choice)
	return c.NextFactor
}

func (c *CorporateActionRate27) AddReinvestmentDiscountRateToMarket() *RateFormat6Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat6Choice)
	return c.ReinvestmentDiscountRateToMarket
}

func (c *CorporateActionRate27) AddInterestShortfall() *RateAndAmountFormat12Choice {
	c.InterestShortfall = new(RateAndAmountFormat12Choice)
	return c.InterestShortfall
}

func (c *CorporateActionRate27) AddRealisedLoss() *RateAndAmountFormat12Choice {
	c.RealisedLoss = new(RateAndAmountFormat12Choice)
	return c.RealisedLoss
}
