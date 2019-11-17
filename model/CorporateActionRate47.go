package model

// Specifies rates related to a corporate action option.
type CorporateActionRate47 struct {

	// Rate proposed in a remarketing of variable rate notes.
	ProposedRate *PercentageRate `xml:"PropsdRate,omitempty"`

	// Rate of allowed over-subscription.
	OversubscriptionRate *RateAndAmountFormat5Choice `xml:"OvrsbcptRate,omitempty"`

	// Requested tax rate in case of breakdown of tax rate, for example, used for adjustment of tax rate. This is the new requested applicable rate.
	RequestedTaxationRate []*RateAndAmountFormat21Choice `xml:"ReqdTaxtnRate,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction in which the income was originally paid, for which relief at source and/or reclaim may be possible.
	RequestedWithholdingOfForeignTax []*RateAndAmountFormat21Choice `xml:"ReqdWhldgOfFrgnTax,omitempty"`

	// Requested rate at which the income will be withheld by the jurisdiction in which the account owner is located, for which relief at source and/or reclaim may be possible.
	RequestedWithholdingOfLocalTax []*RateAndAmountFormat21Choice `xml:"ReqdWhldgOfLclTax,omitempty"`
}

func (c *CorporateActionRate47) SetProposedRate(value string) {
	c.ProposedRate = (*PercentageRate)(&value)
}

func (c *CorporateActionRate47) AddOversubscriptionRate() *RateAndAmountFormat5Choice {
	c.OversubscriptionRate = new(RateAndAmountFormat5Choice)
	return c.OversubscriptionRate
}

func (c *CorporateActionRate47) AddRequestedTaxationRate() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	c.RequestedTaxationRate = append(c.RequestedTaxationRate, newValue)
	return newValue
}

func (c *CorporateActionRate47) AddRequestedWithholdingOfForeignTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	c.RequestedWithholdingOfForeignTax = append(c.RequestedWithholdingOfForeignTax, newValue)
	return newValue
}

func (c *CorporateActionRate47) AddRequestedWithholdingOfLocalTax() *RateAndAmountFormat21Choice {
	newValue := new(RateAndAmountFormat21Choice)
	c.RequestedWithholdingOfLocalTax = append(c.RequestedWithholdingOfLocalTax, newValue)
	return newValue
}
