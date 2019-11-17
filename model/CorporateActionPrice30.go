package model

// Specifies prices.
type CorporateActionPrice30 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat5Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice30) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice30) AddOverSubscriptionDepositPrice() *PriceFormat5Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat5Choice)
	return c.OverSubscriptionDepositPrice
}
