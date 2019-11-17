package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice16 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat19Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice16) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice16) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	newValue := new(PriceFormat20Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice16) AddOverSubscriptionDepositPrice() *PriceFormat19Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat19Choice)
	return c.OverSubscriptionDepositPrice
}
