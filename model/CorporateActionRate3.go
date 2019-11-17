package model

// Specifies rates of a corporate action.
type CorporateActionRate3 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat3Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event. This can be a number or the term "any and all".
	PercentageSought *RateFormat5Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat2Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat2Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat3Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat3Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat3Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat2Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`
}

func (c *CorporateActionRate3) AddInterest() *RateAndAmountFormat3Choice {
	c.Interest = new(RateAndAmountFormat3Choice)
	return c.Interest
}

func (c *CorporateActionRate3) AddPercentageSought() *RateFormat5Choice {
	c.PercentageSought = new(RateFormat5Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate3) AddRelatedIndex() *RateFormat2Choice {
	c.RelatedIndex = new(RateFormat2Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate3) AddSpread() *RateFormat2Choice {
	c.Spread = new(RateFormat2Choice)
	return c.Spread
}

func (c *CorporateActionRate3) AddBidInterval() *RateAndAmountFormat3Choice {
	c.BidInterval = new(RateAndAmountFormat3Choice)
	return c.BidInterval
}

func (c *CorporateActionRate3) AddPreviousFactor() *RateFormat3Choice {
	c.PreviousFactor = new(RateFormat3Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate3) AddNextFactor() *RateFormat3Choice {
	c.NextFactor = new(RateFormat3Choice)
	return c.NextFactor
}

func (c *CorporateActionRate3) AddReinvestmentDiscountRateToMarket() *RateFormat2Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat2Choice)
	return c.ReinvestmentDiscountRateToMarket
}
