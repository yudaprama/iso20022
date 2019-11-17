package model

// Environment common to a collection of transactions.
type CardPaymentEnvironment51 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer5 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation9 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment51) AddAcquirer() *Acquirer5 {
	c.Acquirer = new(Acquirer5)
	return c.Acquirer
}

func (c *CardPaymentEnvironment51) AddMerchant() *Organisation9 {
	c.Merchant = new(Organisation9)
	return c.Merchant
}

func (c *CardPaymentEnvironment51) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}
