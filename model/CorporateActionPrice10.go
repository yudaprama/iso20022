package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice10 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice1Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat11Choice `xml:"CshInLieuOfShrPric,omitempty"`
}

func (c *CorporateActionPrice10) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice1Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice1Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice10) AddCashInLieuOfSharePrice() *PriceFormat11Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat11Choice)
	return c.CashInLieuOfSharePrice
}
