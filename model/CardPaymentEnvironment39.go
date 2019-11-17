package model

// Environment common to a collection of transactions.
type CardPaymentEnvironment39 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer5 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation9 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction4 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment39) AddAcquirer() *Acquirer5 {
	c.Acquirer = new(Acquirer5)
	return c.Acquirer
}

func (c *CardPaymentEnvironment39) AddMerchant() *Organisation9 {
	c.Merchant = new(Organisation9)
	return c.Merchant
}

func (c *CardPaymentEnvironment39) AddPOI() *PointOfInteraction4 {
	c.POI = new(PointOfInteraction4)
	return c.POI
}
