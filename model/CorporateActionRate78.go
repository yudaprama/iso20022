package model

// Specifies rates of a corporate action.
type CorporateActionRate78 struct {

	// Annual rate of a financial instrument.
	Interest *RateAndAmountFormat46Choice `xml:"Intrst,omitempty"`

	// Percentage of securities the offeror/issuer will purchase or redeem under the terms of the event.
	PercentageSought *RateFormat7Choice `xml:"PctgSght,omitempty"`

	// Index rate related to the interest rate of the forthcoming interest payment.
	RelatedIndex *RateFormat3Choice `xml:"RltdIndx,omitempty"`

	// Margin allowed over or under a given rate.
	Spread *RateFormat3Choice `xml:"Sprd,omitempty"`

	// Acceptable price increment used for submitting a bid.
	BidInterval *RateAndAmountFormat50Choice `xml:"BidIntrvl,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) until the next redemption (factor) date.
	PreviousFactor *RateFormat12Choice `xml:"PrvsFctr,omitempty"`

	// Factor used to calculate the value of the outstanding principal of the financial instrument (for factored securities) that will applicable after the redemption (factor) date.
	NextFactor *RateFormat12Choice `xml:"NxtFctr,omitempty"`

	// Rate of discount for securities purchased through a reinvestment scheme as compared to the current market price of security.
	ReinvestmentDiscountRateToMarket *RateFormat3Choice `xml:"RinvstmtDscntRateToMkt,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the interest payment and the expected or scheduled rate of the interest payment .
	InterestShortfall *RateAndAmountFormat43Choice `xml:"IntrstShrtfll,omitempty"`

	// For structured security issues where there is a set schedule of principal and interest payments for the life of the issue, this is the difference between the actual rate of the capital or principal repayment and the scheduled capital repayment.
	RealisedLoss *RateAndAmountFormat43Choice `xml:"RealsdLoss,omitempty"`

	// Dividend or interest rate declared by the issuer.
	DeclaredRate *RateAndAmountFormat43Choice `xml:"DclrdRate,omitempty"`
}

func (c *CorporateActionRate78) AddInterest() *RateAndAmountFormat46Choice {
	c.Interest = new(RateAndAmountFormat46Choice)
	return c.Interest
}

func (c *CorporateActionRate78) AddPercentageSought() *RateFormat7Choice {
	c.PercentageSought = new(RateFormat7Choice)
	return c.PercentageSought
}

func (c *CorporateActionRate78) AddRelatedIndex() *RateFormat3Choice {
	c.RelatedIndex = new(RateFormat3Choice)
	return c.RelatedIndex
}

func (c *CorporateActionRate78) AddSpread() *RateFormat3Choice {
	c.Spread = new(RateFormat3Choice)
	return c.Spread
}

func (c *CorporateActionRate78) AddBidInterval() *RateAndAmountFormat50Choice {
	c.BidInterval = new(RateAndAmountFormat50Choice)
	return c.BidInterval
}

func (c *CorporateActionRate78) AddPreviousFactor() *RateFormat12Choice {
	c.PreviousFactor = new(RateFormat12Choice)
	return c.PreviousFactor
}

func (c *CorporateActionRate78) AddNextFactor() *RateFormat12Choice {
	c.NextFactor = new(RateFormat12Choice)
	return c.NextFactor
}

func (c *CorporateActionRate78) AddReinvestmentDiscountRateToMarket() *RateFormat3Choice {
	c.ReinvestmentDiscountRateToMarket = new(RateFormat3Choice)
	return c.ReinvestmentDiscountRateToMarket
}

func (c *CorporateActionRate78) AddInterestShortfall() *RateAndAmountFormat43Choice {
	c.InterestShortfall = new(RateAndAmountFormat43Choice)
	return c.InterestShortfall
}

func (c *CorporateActionRate78) AddRealisedLoss() *RateAndAmountFormat43Choice {
	c.RealisedLoss = new(RateAndAmountFormat43Choice)
	return c.RealisedLoss
}

func (c *CorporateActionRate78) AddDeclaredRate() *RateAndAmountFormat43Choice {
	c.DeclaredRate = new(RateAndAmountFormat43Choice)
	return c.DeclaredRate
}
