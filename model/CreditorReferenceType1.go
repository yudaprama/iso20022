package model

// Specifies the type of the documents referred by the creditor.
type CreditorReferenceType1 struct {

	// Coded creditor reference type.
	Code *DocumentType3Code `xml:"Cd"`

	// Creditor reference type not avilable in a coded format.
	Proprietary *Max35Text `xml:"Prtry"`

	// Identification of the issuer of the credit reference type.
	Issuer *Max35Text `xml:"Issr,omitempty"`
}

func (c *CreditorReferenceType1) SetCode(value string) {
	c.Code = (*DocumentType3Code)(&value)
}

func (c *CreditorReferenceType1) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}

func (c *CreditorReferenceType1) SetIssuer(value string) {
	c.Issuer = (*Max35Text)(&value)
}
