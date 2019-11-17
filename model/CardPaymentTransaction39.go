package model

// Authorisation response from the acquirer.
type CardPaymentTransaction39 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult4 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult []*TransactionVerificationResult3 `xml:"TxVrfctnRslt,omitempty"`

	// Product code for which the authorisation was declined.
	DeclinedProductCode []*Max70Text `xml:"DclndPdctCd,omitempty"`

	// Balance of the account, related to the payment.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	// Encrypted balance of the account.
	ProtectedBalance *ContentInformationType10 `xml:"PrtctdBal,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action3 `xml:"Actn,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be accepted by the cardholder.
	CurrencyConversion *CurrencyConversion3 `xml:"CcyConvs,omitempty"`
}

func (c *CardPaymentTransaction39) AddAuthorisationResult() *AuthorisationResult4 {
	c.AuthorisationResult = new(AuthorisationResult4)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction39) AddTransactionVerificationResult() *TransactionVerificationResult3 {
	newValue := new(TransactionVerificationResult3)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction39) AddDeclinedProductCode(value string) {
	c.DeclinedProductCode = append(c.DeclinedProductCode, (*Max70Text)(&value))
}

func (c *CardPaymentTransaction39) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardPaymentTransaction39) AddProtectedBalance() *ContentInformationType10 {
	c.ProtectedBalance = new(ContentInformationType10)
	return c.ProtectedBalance
}

func (c *CardPaymentTransaction39) AddAction() *Action3 {
	newValue := new(Action3)
	c.Action = append(c.Action, newValue)
	return newValue
}

func (c *CardPaymentTransaction39) AddCurrencyConversion() *CurrencyConversion3 {
	c.CurrencyConversion = new(CurrencyConversion3)
	return c.CurrencyConversion
}
