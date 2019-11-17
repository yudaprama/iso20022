package model

// Data related to the cardholder.
type Cardholder5 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	Language *ISO2ALanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress13 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress13 `xml:"ShppgAdr,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication5 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder5) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder5) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder5) SetLanguage(value string) {
	c.Language = (*ISO2ALanguageCode)(&value)
}

func (c *Cardholder5) AddBillingAddress() *PostalAddress13 {
	c.BillingAddress = new(PostalAddress13)
	return c.BillingAddress
}

func (c *Cardholder5) AddShippingAddress() *PostalAddress13 {
	c.ShippingAddress = new(PostalAddress13)
	return c.ShippingAddress
}

func (c *Cardholder5) AddAuthentication() *CardholderAuthentication5 {
	newValue := new(CardholderAuthentication5)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder5) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder5) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
