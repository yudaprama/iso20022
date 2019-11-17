package model

// Environment of the transaction.
type CardPaymentEnvironment10 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction2 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard6 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder4 `xml:"Crdhldr,omitempty"`
}

func (c *CardPaymentEnvironment10) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment10) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment10) AddPOI() *PointOfInteraction2 {
	c.POI = new(PointOfInteraction2)
	return c.POI
}

func (c *CardPaymentEnvironment10) AddCard() *PaymentCard6 {
	c.Card = new(PaymentCard6)
	return c.Card
}

func (c *CardPaymentEnvironment10) AddCardholder() *Cardholder4 {
	c.Cardholder = new(Cardholder4)
	return c.Cardholder
}
