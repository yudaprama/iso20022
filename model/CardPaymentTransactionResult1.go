package model

// Result of the transaction.
type CardPaymentTransactionResult1 struct {

	// Type of party that has delivered or declined the card payment authorisation (the party is not identified).
	AuthorisationEntity *GenericIdentification33 `xml:"AuthstnNtty,omitempty"`

	// Response to an authorisation from the acquirer.
	ResponseToAuthorisation *ResponseType1 `xml:"RspnToAuthstn"`

	// Value assigned by the authorising party.
	AuthorisationCode *Min6Max8Text `xml:"AuthstnCd,omitempty"`
}

func (c *CardPaymentTransactionResult1) AddAuthorisationEntity() *GenericIdentification33 {
	c.AuthorisationEntity = new(GenericIdentification33)
	return c.AuthorisationEntity
}

func (c *CardPaymentTransactionResult1) AddResponseToAuthorisation() *ResponseType1 {
	c.ResponseToAuthorisation = new(ResponseType1)
	return c.ResponseToAuthorisation
}

func (c *CardPaymentTransactionResult1) SetAuthorisationCode(value string) {
	c.AuthorisationCode = (*Min6Max8Text)(&value)
}
