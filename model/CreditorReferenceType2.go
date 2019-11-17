package model

// Specifies the type of creditor reference.
type CreditorReferenceType2 struct {

	// Coded or proprietary format creditor reference type.
	CodeOrProprietary *CreditorReferenceType1Choice `xml:"CdOrPrtry"`

	// Entity that assigns the credit reference type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (c *CreditorReferenceType2) AddCodeOrProprietary() *CreditorReferenceType1Choice {
	c.CodeOrProprietary = new(CreditorReferenceType1Choice)
	return c.CodeOrProprietary
}

func (c *CreditorReferenceType2) SetIssuer(value string) {
	c.Issuer = (*Max35Text)(&value)
}
