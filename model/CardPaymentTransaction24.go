package model

// Authorisation response from the acquirer.
type CardPaymentTransaction24 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult2 `xml:"TxVrfctnRslt,omitempty"`

	// Balance of the account, related to the payment.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action3 `xml:"Actn,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
	CurrencyConversion *CurrencyConversion1 `xml:"CcyConvs,omitempty"`
}

func (c *CardPaymentTransaction24) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction24) AddTransactionVerificationResult() *TransactionVerificationResult2 {
	c.TransactionVerificationResult = new(TransactionVerificationResult2)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction24) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardPaymentTransaction24) AddAction() *Action3 {
	newValue := new(Action3)
	c.Action = append(c.Action, newValue)
	return newValue
}

func (c *CardPaymentTransaction24) AddCurrencyConversion() *CurrencyConversion1 {
	c.CurrencyConversion = new(CurrencyConversion1)
	return c.CurrencyConversion
}
