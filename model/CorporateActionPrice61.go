package model

// Specifies prices.
type CorporateActionPrice61 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat50Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat50Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice61) AddCashInLieuOfSharePrice() *PriceFormat50Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat50Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice61) AddOverSubscriptionDepositPrice() *PriceFormat50Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat50Choice)
	return c.OverSubscriptionDepositPrice
}
