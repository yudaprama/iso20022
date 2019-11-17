package model

// Specifies prices of a corporate action.
type CorporateActionPrice17 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat19Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat19Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice17) AddMaximumPrice() *PriceFormat19Choice {
	c.MaximumPrice = new(PriceFormat19Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice17) AddMinimumPrice() *PriceFormat19Choice {
	c.MinimumPrice = new(PriceFormat19Choice)
	return c.MinimumPrice
}
