package model

// Specifies prices of a corporate action.
type CorporateActionPrice57 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat44Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat44Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice57) AddMaximumPrice() *PriceFormat44Choice {
	c.MaximumPrice = new(PriceFormat44Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice57) AddMinimumPrice() *PriceFormat44Choice {
	c.MinimumPrice = new(PriceFormat44Choice)
	return c.MinimumPrice
}
