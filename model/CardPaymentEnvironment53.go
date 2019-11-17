package model

// Environment of the transaction.
type CardPaymentEnvironment53 struct {

	// Acquirer involved in the card payment.
	Acquirer *Acquirer4 `xml:"Acqrr,omitempty"`

	// Merchant performing the card payment transaction.
	// Usage: In some cases, merchant and acceptor may be regarded as the same entity.
	Merchant *Organisation25 `xml:"Mrchnt,omitempty"`

	// Point of interaction (POI) performing the transaction.
	POI *PointOfInteraction5 `xml:"POI,omitempty"`

	// Payment card performing the transaction.
	Card *PaymentCard21 `xml:"Card"`

	// Device used by the customer to perform the payment transaction.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container for tenders used by the customer to perform the payment transaction.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken1 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card payment.
	Cardholder *Cardholder10 `xml:"Crdhldr,omitempty"`

	// Replacement of the message element Cardholder by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardPaymentEnvironment53) AddAcquirer() *Acquirer4 {
	c.Acquirer = new(Acquirer4)
	return c.Acquirer
}

func (c *CardPaymentEnvironment53) AddMerchant() *Organisation25 {
	c.Merchant = new(Organisation25)
	return c.Merchant
}

func (c *CardPaymentEnvironment53) AddPOI() *PointOfInteraction5 {
	c.POI = new(PointOfInteraction5)
	return c.POI
}

func (c *CardPaymentEnvironment53) AddCard() *PaymentCard21 {
	c.Card = new(PaymentCard21)
	return c.Card
}

func (c *CardPaymentEnvironment53) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardPaymentEnvironment53) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardPaymentEnvironment53) AddPaymentToken() *CardPaymentToken1 {
	c.PaymentToken = new(CardPaymentToken1)
	return c.PaymentToken
}

func (c *CardPaymentEnvironment53) AddCardholder() *Cardholder10 {
	c.Cardholder = new(Cardholder10)
	return c.Cardholder
}

func (c *CardPaymentEnvironment53) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}
