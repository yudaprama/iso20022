package model

// Specifies prices.
type CorporateActionPrice63 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat52Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat52Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice63) AddCashInLieuOfSharePrice() *PriceFormat52Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat52Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice63) AddOverSubscriptionDepositPrice() *PriceFormat52Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat52Choice)
	return c.OverSubscriptionDepositPrice
}
