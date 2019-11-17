package model

// Data related to the cardholder.
type Cardholder2 struct {

	// Identification of the cardholder involved in a transaction.
	Identification []*CardholderIdentification1 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Data related to the authentication of the cardholder.
	Authentication []*CardholderAuthentication2 `xml:"Authntcn,omitempty"`

	// Numeric characters of the cardholder's address for verification.
	AddressVerification *AddressVerification1 `xml:"AdrVrfctn,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder2) AddIdentification() *CardholderIdentification1 {
	newValue := new(CardholderIdentification1)
	c.Identification = append(c.Identification, newValue)
	return newValue
}

func (c *Cardholder2) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder2) AddAuthentication() *CardholderAuthentication2 {
	newValue := new(CardholderAuthentication2)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder2) AddAddressVerification() *AddressVerification1 {
	c.AddressVerification = new(AddressVerification1)
	return c.AddressVerification
}

func (c *Cardholder2) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
