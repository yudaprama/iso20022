package model

// Specifies prices.
type CorporateActionPrice64 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat52Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice9Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice4 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat55Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat56Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice64) AddCashInLieuOfSharePrice() *PriceFormat52Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat52Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice64) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice9Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice9Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice64) AddCashValueForTax() *AmountPrice4 {
	c.CashValueForTax = new(AmountPrice4)
	return c.CashValueForTax
}

func (c *CorporateActionPrice64) AddGenericCashPricePaidPerProduct() *PriceFormat55Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat55Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice64) AddGenericCashPriceReceivedPerProduct() *PriceFormat56Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat56Choice)
	return c.GenericCashPriceReceivedPerProduct
}
