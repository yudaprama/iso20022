package model

// Environment of the transaction.
type CardPaymentEnvironment20 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer2 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation8 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction3 `xml:"POI"`

	// Payment card performing the transaction.
	Card *PaymentCard7 `xml:"Card"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder5 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType7 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment20) AddAcquirer() *Acquirer2 {
	c.Acquirer = new(Acquirer2)
	return c.Acquirer
}

func (c *CardPaymentEnvironment20) AddMerchant() *Organisation8 {
	c.Merchant = new(Organisation8)
	return c.Merchant
}

func (c *CardPaymentEnvironment20) AddPOI() *PointOfInteraction3 {
	c.POI = new(PointOfInteraction3)
	return c.POI
}

func (c *CardPaymentEnvironment20) AddCard() *PaymentCard7 {
	c.Card = new(PaymentCard7)
	return c.Card
}

func (c *CardPaymentEnvironment20) AddCardholder() *Cardholder5 {
	c.Cardholder = new(Cardholder5)
	return c.Cardholder
}

func (c *CardPaymentEnvironment20) AddProtectedCardholderData() *ContentInformationType7 {
	c.ProtectedCardholderData = new(ContentInformationType7)
	return c.ProtectedCardholderData
}
