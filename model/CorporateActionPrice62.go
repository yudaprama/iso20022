package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice62 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice9Choice `xml:"IndctvOrMktPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat52Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat53Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat52Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice62) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice9Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice9Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice62) AddIssuePrice() *PriceFormat52Choice {
	c.IssuePrice = new(PriceFormat52Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice62) AddGenericCashPriceReceivedPerProduct() *PriceFormat53Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat53Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice62) AddGenericCashPricePaidPerProduct() *PriceFormat52Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat52Choice)
	return c.GenericCashPricePaidPerProduct
}
