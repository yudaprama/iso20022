package model

// Authorisation response from the acquirer.
type CardPaymentTransaction54 struct {

	// Outcome of the authorisation, and actions to perform.
	AuthorisationResult *AuthorisationResult10 `xml:"AuthstnRslt"`

	// Result of the verifications performed by the issuer to deliver or decline the authorisation.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Product code which are allowed by the payment card.
	AllowedProductCode []*Product4 `xml:"AllwdPdctCd,omitempty"`

	// Product code not allowed by the payment card.
	NotAllowedProductCode []*Product4 `xml:"NotAllwdPdctCd,omitempty"`

	// Products that may be added to the purchase after the authorisation.
	AdditionalAvailableProduct []*Product5 `xml:"AddtlAvlblPdct,omitempty"`

	// Balance of the account, related to the payment.
	Balance *AmountAndDirection41 `xml:"Bal,omitempty"`

	// Encrypted balance of the account.
	ProtectedBalance *ContentInformationType10 `xml:"PrtctdBal,omitempty"`

	// Set of actions to be performed by the POI (Point Of Interaction) system.
	Action []*Action6 `xml:"Actn,omitempty"`

	// Conversion between the currency of a card acceptor and the currency of a card issuer, provided by a dedicated service provider. The currency conversion has to be proposed to the cardholder.
	CurrencyConversionEligibility *CurrencyConversion6 `xml:"CcyConvsElgblty,omitempty"`
}

func (c *CardPaymentTransaction54) AddAuthorisationResult() *AuthorisationResult10 {
	c.AuthorisationResult = new(AuthorisationResult10)
	return c.AuthorisationResult
}

func (c *CardPaymentTransaction54) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddAllowedProductCode() *Product4 {
	newValue := new(Product4)
	c.AllowedProductCode = append(c.AllowedProductCode, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddNotAllowedProductCode() *Product4 {
	newValue := new(Product4)
	c.NotAllowedProductCode = append(c.NotAllowedProductCode, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddAdditionalAvailableProduct() *Product5 {
	newValue := new(Product5)
	c.AdditionalAvailableProduct = append(c.AdditionalAvailableProduct, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddBalance() *AmountAndDirection41 {
	c.Balance = new(AmountAndDirection41)
	return c.Balance
}

func (c *CardPaymentTransaction54) AddProtectedBalance() *ContentInformationType10 {
	c.ProtectedBalance = new(ContentInformationType10)
	return c.ProtectedBalance
}

func (c *CardPaymentTransaction54) AddAction() *Action6 {
	newValue := new(Action6)
	c.Action = append(c.Action, newValue)
	return newValue
}

func (c *CardPaymentTransaction54) AddCurrencyConversionEligibility() *CurrencyConversion6 {
	c.CurrencyConversionEligibility = new(CurrencyConversion6)
	return c.CurrencyConversionEligibility
}
