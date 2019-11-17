package model

// Specifies prices related to a corporate action option.
type CorporateActionPrice18 struct {

	// Indicates whether the price is an indicative price or a market price.
	IndicativeOrMarketPrice *IndicativeOrMarketPrice5Choice `xml:"IndctvOrMktPric,omitempty"`

	// Cash disbursement in lieu of equities; usually in lieu of fractional quantity.
	CashInLieuOfSharePrice *PriceFormat19Choice `xml:"CshInLieuOfShrPric,omitempty"`
}

func (c *CorporateActionPrice18) AddIndicativeOrMarketPrice() *IndicativeOrMarketPrice5Choice {
	c.IndicativeOrMarketPrice = new(IndicativeOrMarketPrice5Choice)
	return c.IndicativeOrMarketPrice
}

func (c *CorporateActionPrice18) AddCashInLieuOfSharePrice() *PriceFormat19Choice {
	c.CashInLieuOfSharePrice = new(PriceFormat19Choice)
	return c.CashInLieuOfSharePrice
}
