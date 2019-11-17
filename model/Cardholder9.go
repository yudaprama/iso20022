package model

// Data related to the cardholder.
type Cardholder9 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress18 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress18 `xml:"ShppgAdr,omitempty"`

	// Method and data intended to be used for this transaction to authenticate the cardholder and its card.
	Authentication []*CardholderAuthentication7 `xml:"Authntcn,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder9) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder9) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder9) SetLanguage(value string) {
	c.Language = (*LanguageCode)(&value)
}

func (c *Cardholder9) AddBillingAddress() *PostalAddress18 {
	c.BillingAddress = new(PostalAddress18)
	return c.BillingAddress
}

func (c *Cardholder9) AddShippingAddress() *PostalAddress18 {
	c.ShippingAddress = new(PostalAddress18)
	return c.ShippingAddress
}

func (c *Cardholder9) AddAuthentication() *CardholderAuthentication7 {
	newValue := new(CardholderAuthentication7)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder9) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *Cardholder9) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
