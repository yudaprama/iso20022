package model

// Specifies prices of a corporate action.
type CorporateActionPrice67 struct {

	// Maximum or cap price at which a holder can bid, for example, on a Dutch auction offer.
	MaximumPrice *PriceFormat59Choice `xml:"MaxPric,omitempty"`

	// Minimum or floor price at which a holder can bid, for example, on a Dutch auction offer.
	MinimumPrice *PriceFormat59Choice `xml:"MinPric,omitempty"`
}

func (c *CorporateActionPrice67) AddMaximumPrice() *PriceFormat59Choice {
	c.MaximumPrice = new(PriceFormat59Choice)
	return c.MaximumPrice
}

func (c *CorporateActionPrice67) AddMinimumPrice() *PriceFormat59Choice {
	c.MinimumPrice = new(PriceFormat59Choice)
	return c.MinimumPrice
}
