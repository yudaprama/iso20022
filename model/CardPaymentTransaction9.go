package model

// Authorisation response from the acquirer.
type CardPaymentTransaction9 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult1 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult *TransactionVerificationResult1 `xml:"TxVrfctnRslt,omitempty"`

	// Balance of the account, related to the payment.
	Balance *ImpliedCurrencyAndAmount `xml:"Bal,omitempty"`

	// Currency associated with the transaction.
	Currency *CurrencyCode `xml:"Ccy,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action1 `xml:"Actn,omitempty"`
}

func (c *CardPaymentTransaction9) AddAuthorisationResult() *AuthorisationResult1 {
	c.AuthorisationResult = new(AuthorisationResult1)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction9) AddTransactionVerificationResult() *TransactionVerificationResult1 {
	c.TransactionVerificationResult = new(TransactionVerificationResult1)
	return c.TransactionVerificationResult
}

func (c *CardPaymentTransaction9) SetBalance(value, currency string) {
	c.Balance = NewImpliedCurrencyAndAmount(value, currency)
}

func (c *CardPaymentTransaction9) SetCurrency(value string) {
	c.Currency = (*CurrencyCode)(&value)
}

func (c *CardPaymentTransaction9) AddAction() *Action1 {
	newValue := new(Action1)
	c.Action = append(c.Action, newValue)
	return newValue
}
