package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice65 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat57Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat57Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice65) AddCashInLieuOfSharePrice() *PriceFormat57Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat57Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice65) AddOverSubscriptionDepositPrice() *PriceFormat57Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat57Choice)
	return c.OverSubscriptionDepositPrice
}
