package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice38 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice5Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat29Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat23Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat20Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice38) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice5Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice5Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice38) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice38) AddCashValueForTax() *PriceFormat29Choice {
	c.CashValueForTax = new(PriceFormat29Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice38) AddGenericCashPricePaidPerProduct() *PriceFormat23Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat23Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice38) AddGenericCashPriceReceivedPerProduct() *PriceFormat20Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat20Choice)
	return c.GenericCashPriceReceivedPerProduct
}
