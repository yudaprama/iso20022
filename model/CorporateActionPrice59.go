package model

// Specifies prices.
type CorporateActionPrice59 struct {

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat50Choice `xml:"CshInLieuOfShrPric,omitempty"`

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice8Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash value of resulting securities proceeds for tax calculation and/or reporting.
	CashValueForTax *AmountPrice2 `xml:"CshValForTax,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount or a number of points above an index, for example, reinvestment price, strike price and exercise price.
	GenericCashPricePaidPerProduct *PriceFormat51Choice `xml:"GncCshPricPdPerPdct,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat48Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`
}

func (c *CorporateActionPrice59) AddCashInLieuOfSharePrice() *PriceFormat50Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat50Choice)
	return c.CashInLieuOfSharePrice
}

func (c *CorporateActionPrice59) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice8Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice8Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice59) AddCashValueForTax() *AmountPrice2 {
	c.CashValueForTax = new(AmountPrice2)
	return c.CashValueForTax
}

func (c *CorporateActionPrice59) AddGenericCashPricePaidPerProduct() *PriceFormat51Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat51Choice)
	return c.GenericCashPricePaidPerProduct
}

func (c *CorporateActionPrice59) AddGenericCashPriceReceivedPerProduct() *PriceFormat48Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat48Choice)
	return c.GenericCashPriceReceivedPerProduct
}
