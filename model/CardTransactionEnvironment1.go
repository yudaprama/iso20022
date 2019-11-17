package model

// Environment of the transaction.
type CardTransactionEnvironment1 struct {

	// Acquirer of the card transaction.
	Acquirer *Acquirer6 `xml:"Acqrr"`

	// Identification of the interconnected card scheme from which the request is coming.
	CardSchemeIdentification *Max35Text `xml:"CardSchmeId,omitempty"`

	// Acceptor performing the card transaction.
	Acceptor *Organisation18 `xml:"Accptr,omitempty"`

	// Payment terminal or ATM performing the transaction.
	Terminal *CardAcceptorTerminal1 `xml:"Termnl,omitempty"`

	// Card performing the transaction.
	Card *PaymentCard12 `xml:"Card"`

	// Container of tenders used by the customer to perform the payment.
	CustomerDevice *CustomerDevice1 `xml:"CstmrDvc,omitempty"`

	// Container of tenders used by the customer to perform the payment.
	Wallet *CustomerDevice1 `xml:"Wllt,omitempty"`

	// Payment token information.
	PaymentToken *CardPaymentToken4 `xml:"PmtTkn,omitempty"`

	// Cardholder involved in the card transaction.
	// It correspond partially to the ISO 8583:2003, field number 49-71.
	Cardholder *Cardholder9 `xml:"Crdhldr,omitempty"`

	// Protection of cardholder sensitive data by a digital envelope using a cryptographic key.
	ProtectedCardholderData *ContentInformationType10 `xml:"PrtctdCrdhldrData,omitempty"`
}

func (c *CardTransactionEnvironment1) AddAcquirer() *Acquirer6 {
	c.Acquirer = new(Acquirer6)
	return c.Acquirer
}

func (c *CardTransactionEnvironment1) SetCardSchemeIdentification(value string) {
	c.CardSchemeIdentification = (*Max35Text)(&value)
}

func (c *CardTransactionEnvironment1) AddAcceptor() *Organisation18 {
	c.Acceptor = new(Organisation18)
	return c.Acceptor
}

func (c *CardTransactionEnvironment1) AddTerminal() *CardAcceptorTerminal1 {
	c.Terminal = new(CardAcceptorTerminal1)
	return c.Terminal
}

func (c *CardTransactionEnvironment1) AddCard() *PaymentCard12 {
	c.Card = new(PaymentCard12)
	return c.Card
}

func (c *CardTransactionEnvironment1) AddCustomerDevice() *CustomerDevice1 {
	c.CustomerDevice = new(CustomerDevice1)
	return c.CustomerDevice
}

func (c *CardTransactionEnvironment1) AddWallet() *CustomerDevice1 {
	c.Wallet = new(CustomerDevice1)
	return c.Wallet
}

func (c *CardTransactionEnvironment1) AddPaymentToken() *CardPaymentToken4 {
	c.PaymentToken = new(CardPaymentToken4)
	return c.PaymentToken
}

func (c *CardTransactionEnvironment1) AddCardholder() *Cardholder9 {
	c.Cardholder = new(Cardholder9)
	return c.Cardholder
}

func (c *CardTransactionEnvironment1) AddProtectedCardholderData() *ContentInformationType10 {
	c.ProtectedCardholderData = new(ContentInformationType10)
	return c.ProtectedCardholderData
}
