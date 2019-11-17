package model

// Specifies rates related to a corporate action option.
type CorporateActionRate71 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat39Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	RequestedWithholdingTaxRate []*RateAndAmountFormat40Choice `xml:"ReqdWhldgTaxRate,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible.
	RequestedSecondLevelTaxRate []*RateAndAmountFormat40Choice `xml:"ReqdScndLvlTaxRate,omitempty"`
}

func (c *CorporateActionRate71) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate71) AddOversubscriptionRate() *RateAndAmountFormat39Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat39Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate71) AddRequestedWithholdingTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.RequestedWithholdingTaxRate = append(c.RequestedWithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate71) AddRequestedSecondLevelTaxRate() *RateAndAmountFormat40Choice {
	newValue := new(RateAndAmountFormat40Choice)
	c.RequestedSecondLevelTaxRate = append(c.RequestedSecondLevelTaxRate, newValue)
	return newValue
}
