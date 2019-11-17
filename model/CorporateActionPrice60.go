package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice60 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice8Choice `xml:"IndctvOrMktPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat50Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat49Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat50Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice60) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice8Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice8Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice60) AddIssuePrice() *PriceFormat50Choice {
	c.IssuePrice = new(PriceFormat50Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice60) AddGenericCashPriceReceivedPerProduct() *PriceFormat49Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat49Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice60) AddGenericCashPricePaidPerProduct() *PriceFormat50Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat50Choice)
	return c.GenericCashPricePaidPerProduct
}
