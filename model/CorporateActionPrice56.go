package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice56 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice7Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat45Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat46Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat44Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat47Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice56) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice7Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice7Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice56) AddCashInLieuOfSharePrice() *PriceFormat45Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat45Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice56) AddCashValueForTax() *PriceFormat46Choice {
	c.CashValueForTax = new(PriceFormat46Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice56) AddGenericCashPricePaidPerProduct() *PriceFormat44Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat44Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice56) AddGenericCashPriceReceivedPerProduct() *PriceFormat47Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat47Choice)
	return c.GenericCashPriceReceivedPerProduct
}
