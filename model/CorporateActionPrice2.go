package model

// Specifies prices.
type CorporateActionPrice2 struct {

	// Maximum or cap price at which a holder can bid, e.g. on a Dutch auction offer.
	MaximumPrice *PriceFormat3Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, e.g. on a Dutch auction offer.
	MinimumPrice *PriceFormat3Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice2) AddMaximumPrice() *PriceFormat3Choice {
	c.MaximumPrice = new(PriceFormat3Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice2) AddMinimumPrice() *PriceFormat3Choice {
	c.MinimumPrice = new(PriceFormat3Choice)
	return c.MinimumPrice
}
