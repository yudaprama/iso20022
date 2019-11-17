package model

// Data related to the cardholder.
type Cardholder1 struct {

	// Identification of the cardholder involved in a transaction.
	Identification []*CardholderIdentification1 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	Language *ISO2ALanguageCode `xml:"Lang,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication1 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder1) AddIdentification() *CardholderIdentification1 {
	newValue := new(CardholderIdentification1)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *Cardholder1) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder1) SetLanguage(value string) {
	c.Language = (*ISO2ALanguageCode)(&value)
}

func (c *Cardholder1) AddAuthentication() *CardholderAuthentication1 {
	newValue := new(CardholderAuthentication1)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder1) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder1) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
