package model

// Specifies rates related to a corporate action option.
type CorporateActionRate8 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat12Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate in case of breakdown of tax rate, for example, used for adjustment of tax rate. This is the new requested applicable rate.
	RequestedTaxationRate *PercentageRate `xml:"ReqdTaxtnRate,omitempty"`
}

func (c *CorporateActionRate8) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate8) AddOversubscriptionRate() *RateAndAmountFormat12Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat12Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate8) SetRequestedTaxationRate(value string) {
	c.RequestedTaxationRate = (*PercentageRate)(&value)
}
