package model

// Specifies prices.
type CorporateActionPrice21 struct {

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct []*PriceFormat7Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat5Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Amount of money required per over-subscribed equity as defined by the issuer.
	OverSubscriptionDepositPrice *PriceFormat5Choice `xml:"OverSbcptDpstPric,omitempty"`
}

func (c *CorporateActionPrice21) AddGenericCashPriceReceivedPerProduct() *PriceFormat7Choice {
	newValue := new(PriceFormat7Choice)
	c.GenericCashPriceReceivedPerProduct = append(c.GenericCashPriceReceivedPerProduct, newValue)
	return newValue
}

func (c *CorporateActionPrice21) AddCashInLieuOfSharePrice() *PriceFormat5Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat5Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice21) AddOverSubscriptionDepositPrice() *PriceFormat5Choice {
	c.OverSubscriptionDepositPrice = new(PriceFormat5Choice)
	return c.OverSubscriptionDepositPrice
}
