package model

// Data related to the cardholder.
type Cardholder7 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress13 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress13 `xml:"ShppgAdr,omitempty"`

	// Method and data intended to be used for this transaction to authenticate the cardholder.
	Authentication []*CardholderAuthentication6 `xml:"Authntcn,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult3 `xml:"TxVrfctnRslt,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder7) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder7) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder7) SetLanguage(value string) {
	c.Language = (*LanguageCode)(&value)
}

func (c *Cardholder7) AddBillingAddress() *PostalAddress13 {
	c.BillingAddress = new(PostalAddress13)
	return c.BillingAddress
}

func (c *Cardholder7) AddShippingAddress() *PostalAddress13 {
	c.ShippingAddress = new(PostalAddress13)
	return c.ShippingAddress
}

func (c *Cardholder7) AddAuthentication() *CardholderAuthentication6 {
	newValue := new(CardholderAuthentication6)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder7) AddTransactionVerificationResult() *TransactionVerificationResult3 {
	newValue := new(TransactionVerificationResult3)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *Cardholder7) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
