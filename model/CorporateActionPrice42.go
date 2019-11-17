package model

// Specifies prices of a corporate action.
type CorporateActionPrice42 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat23Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat23Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice42) AddMaximumPrice() *PriceFormat23Choice {
	c.MaximumPrice = new(PriceFormat23Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice42) AddMinimumPrice() *PriceFormat23Choice {
	c.MinimumPrice = new(PriceFormat23Choice)
	return c.MinimumPrice
}
