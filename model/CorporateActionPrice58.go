package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice58 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat45Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat45Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice58) AddCashInLieuOfSharePrice() *PriceFormat45Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat45Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice58) AddOverSubscriptionDepositPrice() *PriceFormat45Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat45Choice)
	return c.OverSubscriptionDepositPrice
}
