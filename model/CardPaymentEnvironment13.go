package model

// Environment common to a collection of transactions.
type CardPaymentEnvironment13 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer3 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation9 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment13) AddAcquirer() *Acquirer3 {
	c.Acquirer = new(Acquirer3)
	return c.Acquirer
}

func (c *CardPaymentEnvironment13) AddMerchant() *Organisation9 {
	c.Merchant = new(Organisation9)
	return c.Merchant
}

func (c *CardPaymentEnvironment13) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}
