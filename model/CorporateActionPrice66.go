package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice66 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice11Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat57Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *PriceFormat58Choice `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat59Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat60Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice66) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice11Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice11Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice66) AddCashInLieuOfSharePrice() *PriceFormat57Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat57Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice66) AddCashValueForTax() *PriceFormat58Choice {
	c.CashValueForTax = new(PriceFormat58Choice)
	return c.CashValueForTax
}

func (c *CorporateActionPrice66) AddGenericCashPricePaidPerProduct() *PriceFormat59Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat59Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice66) AddGenericCashPriceReceivedPerProduct() *PriceFormat60Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat60Choice)
	return c.GenericCashPriceReceivedPerProduct
}
