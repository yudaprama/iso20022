package model

// Environment of the card payment transaction.
type CardPaymentEnvironment52 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard20 `xml:"Card"`

	// Device used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information
	PaymentToken *CardPaymentToken3 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder11 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment52) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment52) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment52) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment52) AddCard() *PaymentCard20 {
	c.Card = new(PaymentCard20)
	return c.Card
}

func (c *CardPaymentEnvironment52) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment52) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment52) AddPaymentToken() *CardPaymentToken3 {
	c.PaymentToken = new(CardPaymentToken3)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment52) AddCardholder() *Cardholder11 {
	c.Cardholder = new(Cardholder11)
	return c.Cardholder
}

func (c *CardPaymentEnvironment52) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}
