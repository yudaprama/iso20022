package model

// Data related to the cardholder.
type Cardholder10 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification11 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Language selected for the cardholder interface during the transaction.
	// Reference ISO 639-1 (alpha-2) et ISO 639-2 (alpha-3).
	Language *LanguageCode `xml:"Lang,omitempty"`

	// Postal address of the owner of the payment card.
	BillingAddress *PostalAddress18 `xml:"BllgAdr,omitempty"`

	// Postal address for delivery of goods or services.
	ShippingAddress *PostalAddress18 `xml:"ShppgAdr,omitempty"`

	// Identification of the trip.
	TripNumber *Max35Text `xml:"TripNb,omitempty"`

	// Information related to the vehicle used for the transaction.
	Vehicle *Vehicle1 `xml:"Vhcl,omitempty"`

	// Method and data intended to be used for this transaction to authenticate the cardholder and its card.
	Authentication []*CardholderAuthentication9 `xml:"Authntcn,omitempty"`

	// Result of performed verifications for the transaction.
	TransactionVerificationResult []*TransactionVerificationResult4 `xml:"TxVrfctnRslt,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder10) AddIdentification() *PersonIdentification11 {
	c.Identification = new(PersonIdentification11)
	return c.Identification
}

func (c *Cardholder10) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder10) SetLanguage(value string) {
	c.Language = (*LanguageCode)(&value)
}

func (c *Cardholder10) AddBillingAddress() *PostalAddress18 {
	c.BillingAddress = new(PostalAddress18)
	return c.BillingAddress
}

func (c *Cardholder10) AddShippingAddress() *PostalAddress18 {
	c.ShippingAddress = new(PostalAddress18)
	return c.ShippingAddress
}

func (c *Cardholder10) SetTripNumber(value string) {
	c.TripNumber = (*Max35Text)(&value)
}

func (c *Cardholder10) AddVehicle() *Vehicle1 {
	c.Vehicle = new(Vehicle1)
	return c.Vehicle
}

func (c *Cardholder10) AddAuthentication() *CardholderAuthentication9 {
	newValue := new(CardholderAuthentication9)
	c.Authentication = append(c.Authentication, newValue)
	return newValue
}

func (c *Cardholder10) AddTransactionVerificationResult() *TransactionVerificationResult4 {
	newValue := new(TransactionVerificationResult4)
	c.TransactionVerificationResult = append(c.TransactionVerificationResult, newValue)
	return newValue
}

func (c *Cardholder10) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
