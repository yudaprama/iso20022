package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice28 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat19Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice28) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice28) AddOverSubscriptionDepositPrice() *PriceFormat19Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat19Choice)
	return c.OverSubscriptionDepositPrice
}
