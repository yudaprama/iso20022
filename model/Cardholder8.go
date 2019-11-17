package model

// Data related to the cardholder.
type Cardholder8 struct {

	// Identification of the cardholder involved in a transaction.
	Identification *PersonIdentification7 `xml:"Id,omitempty"`

	// Cardholder name associated with the card.
	Name *Max45Text `xml:"Nm,omitempty"`

	// Identifies personal data related to the cardholder.
	PersonalData *Max70Text `xml:"PrsnlData,omitempty"`
}

func (c *Cardholder8) AddIdentification() *PersonIdentification7 {
	c.Identification = new(PersonIdentification7)
	return c.Identification
}

func (c *Cardholder8) SetName(value string) {
	c.Name = (*Max45Text)(&value)
}

func (c *Cardholder8) SetPersonalData(value string) {
	c.PersonalData = (*Max70Text)(&value)
}
