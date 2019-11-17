package model

// Identification of the cardholder involved in a transaction.
type CardholderIdentification1 struct {

	// Identification value of the cardholder involved in a transaction.
	CardholderIdentificationValue *Max35Text `xml:"CrdhldrIdVal"`

	// Type of identification used for identifying the cardholder.
	CardholderIdentificationType *PersonIdentificationType4Code `xml:"CrdhldrIdTp"`
}

func (c *CardholderIdentification1) SetCardholderIdentificationValue(value string) {
	c.CardholderIdentificationValue = (*Max35Text)(&value)
}

func (c *CardholderIdentification1) SetCardholderIdentificationType(value string) {
	c.CardholderIdentificationType = (*PersonIdentificationType4Code)(&value)
}
