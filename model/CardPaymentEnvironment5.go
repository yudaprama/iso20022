package model

// Environment common to a collection of transactions.
type CardPaymentEnvironment5 struct {

	// Acquirer involved in the card payment transactions.
	Acquirer *Acquirer1 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transactions.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation5 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction1 `xml:"POI,omitempty"`
}

func (c *CardPaymentEnvironment5) AddAcquirer() *Acquirer1 {
	c.Acquirer = new(Acquirer1)
	return c.Acquirer
}

func (c *CardPaymentEnvironment5) AddMerchant() *Organisation5 {
	c.Merchant = new(Organisation5)
	return c.Merchant
}

func (c *CardPaymentEnvironment5) AddPOI() *PointOfInteraction1 {
	c.POI = new(PointOfInteraction1)
	return c.POI
}
