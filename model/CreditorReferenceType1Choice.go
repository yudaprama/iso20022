package model

// Specifies the type of document referred by the creditor.
type CreditorReferenceType1Choice struct {

	// Type of creditor reference, in a coded form.
	Code *DocumentType3Code `xml:"Cd"`

	// Creditor reference type, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (c *CreditorReferenceType1Choice) SetCode(value string) {
	c.Code = (*DocumentType3Code)(&value)
}

func (c *CreditorReferenceType1Choice) SetProprietary(value string) {
	c.Proprietary = (*Max35Text)(&value)
}
