package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice29 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat5Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat21Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice29) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice29) AddIssuePrice() *PriceFormat5Choice {
	c.IssuePrice = new(PriceFormat5Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice29) AddGenericCashPriceReceivedPerProduct() *PriceFormat21Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat21Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice29) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}
