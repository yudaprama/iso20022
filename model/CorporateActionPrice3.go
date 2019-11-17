package model

// Specifies prices of a corporate action.
type CorporateActionPrice3 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat11Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat11Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice3) AddMaximumPrice() *PriceFormat11Choice {
	c.MaximumPrice = new(PriceFormat11Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice3) AddMinimumPrice() *PriceFormat11Choice {
	c.MinimumPrice = new(PriceFormat11Choice)
	return c.MinimumPrice
}
