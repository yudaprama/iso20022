package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice19 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice2Choice `xml:"IndctvOrMktPric,omitempty"`

	// 1. Price at which security will be purchased/sold if warrant is exercised, either as an actual amount or a percentage.
	// 2. Price at which a bond is converted to underlying security either as an actual amount or a percentage.
	// 3. Strike price of an option, represented either as an actual amount, a percentage or a number of points above an index.
	ExercisePrice *PriceFormat5Choice `xml:"ExrcPric,omitempty"`

	// Initial issue price of a financial instrument.
	IssuePrice *PriceFormat5Choice `xml:"IssePric,omitempty"`

	// Generic cash price received per product by the underlying security holder either as a percentage or an amount, for example, redemption price.
	GenericCashPriceReceivedPerProduct *PriceFormat21Choice `xml:"GncCshPricRcvdPerPdct,omitempty"`

	// Generic cash price paid per product by the underlying security holder either as a percentage or an amount, for example, reinvestment price.
	GenericCashPricePaidPerProduct *PriceFormat5Choice `xml:"GncCshPricPdPerPdct,omitempty"`
}

func (c *CorporateActionPrice19) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice2Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice2Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice19) AddExercisePrice() *PriceFormat5Choice {
	c.ExercisePrice = new(PriceFormat5Choice)
	return c.ExercisePrice
}

func (c *CorporateActionPrice19) AddIssuePrice() *PriceFormat5Choice {
	c.IssuePrice = new(PriceFormat5Choice)
	return c.IssuePrice
}

func (c *CorporateActionPrice19) AddGenericCashPriceReceivedPerProduct() *PriceFormat21Choice {
	c.GenericCashPriceReceivedPerProduct = new(PriceFormat21Choice)
	return c.GenericCashPriceReceivedPerProduct
}

func (c *CorporateActionPrice19) AddGenericCashPricePaidPerProduct() *PriceFormat5Choice {
	c.GenericCashPricePaidPerProduct = new(PriceFormat5Choice)
	return c.GenericCashPricePaidPerProduct
}
