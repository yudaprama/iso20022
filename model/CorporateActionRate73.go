package model

// Specifies rates related to a corporate action option.
type CorporateActionRate73 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat43Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate that will be withheld by the tax authorities of the jurisdiction of the issuer, for which a relief at source and/or reclaim may be possible.
	RequestedWithholdingTaxRate []*RateAndAmountFormat45Choice `xml:"ReqdWhldgTaxRate,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction other than the jurisdiction of the issuer’s country of tax incorporation, for which a relief at source and/or reclaim may be possible.
	RequestedSecondLevelTaxRate []*RateAndAmountFormat45Choice `xml:"ReqdScndLvlTaxRate,omitempty"`
}

func (c *CorporateActionRate73) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate73) AddOversubscriptionRate() *RateAndAmountFormat43Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat43Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate73) AddRequestedWithholdingTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.RequestedWithholdingTaxRate = append(c.RequestedWithholdingTaxRate, newValue)
	return newValue
}

func (c *CorporateActionRate73) AddRequestedSecondLevelTaxRate() *RateAndAmountFormat45Choice {
	newValue := new(RateAndAmountFormat45Choice)
	c.RequestedSecondLevelTaxRate = append(c.RequestedSecondLevelTaxRate, newValue)
	return newValue
}
