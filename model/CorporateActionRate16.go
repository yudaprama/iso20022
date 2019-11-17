package model

// Specifies rates of a corporate action.
type CorporateActionRate16 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat14Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	PercentageSought *RateFormat5Choice `xml:"PctgSght,omitempty"`

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
}

func (c *CorporateActionRate16) AddInterest() *RateAndAmountFormat14Choice {
	c.Interest = new(RateAndAmountFormat14Choice)
	return c.Interest
}

func (c *CorporateActionRate16) AddPercentageSought() *RateFormat5Choice {
	c.PercentageSought = new(RateFormat5Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate16) AddRelatedIndex() *RateFormat6Choice {
	c.RelatedIndex = new(RateFormat6Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate16) AddSpread() *RateFormat6Choice {
	c.Spread = new(RateFormat6Choice)
	return c.Spread
}

func (c *CorporateActionRate16) AddBidInterval() *RateAndAmountFormat14Choice {
	c.BidInterval = new(RateAndAmountFormat14Choice)
	return c.BidInterval
}

func (c *CorporateActionRate16) AddPreviousFactor() *RateFormat3Choice {
	c.PreviousFactor = new(RateFormat3Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate16) AddNextFactor() *RateFormat3Choice {
	c.NextFactor = new(RateFormat3Choice)
	return c.NextFactor
}

func (c *CorporateActionRate16) AddReinvestmentDiscountRateToMarket() *RateFormat6Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat6Choice)
	return c.ReinvestmentDiscountRateToMarket
}
