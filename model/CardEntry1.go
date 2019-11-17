package model

// Card transaction entry.
type CardEntry1 struct {

	// Electronic money product that provides the cardholder with a portable and specialised computer device, which typically contains a microprocessor.
	Card *PaymentCard4 `xml:"Card,omitempty"`

	// Physical or logical card payment terminal containing software and hardware components.
	POI *PointOfInteraction1 `xml:"POI,omitempty"`

	// Card entry details, based on card transaction aggregated data performed by the account servicer.
	AggregatedEntry *CardAggregated1 `xml:"AggtdNtry,omitempty"`
}

func (c *CardEntry1) AddCard() *PaymentCard4 {
	c.Card = new(PaymentCard4)
	return c.Card
}

func (c *CardEntry1) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}

func (c *CardEntry1) AddAggregatedEntry() *CardAggregated1 {
	c.AggregatedEntry = new(CardAggregated1)
	return c.AggregatedEntry
}
