package model

// Environment of the card payment transaction.
type CardPaymentEnvironment27 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard8 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder6 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType7 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment27) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment27) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment27) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment27) AddCard() *PaymentCard8 {
	c.Card = new(PaymentCard8)
	return c.Card
}

func (c *CardPaymentEnvironment27) AddCardholder() *Cardholder6 {
	c.Cardholder = new(Cardholder6)
	return c.Cardholder
}

func (c *CardPaymentEnvironment27) AddProtectedCardholderData() *ContentInformationType7 {
	c.ProtectedCardholderData = new(ContentInformationType7)
	return c.ProtectedCardholderData
}
