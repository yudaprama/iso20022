package model

// Choice of formats for the type of mail.
type MailType1Choice struct {

	// Mail type expressed as a code.
	Code *MailType1Code `xml:"Cd"`

	// Mail type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (m *MailType1Choice) SetCode(value string) {
	m.Code = (*MailType1Code)(&value)
}

func (m *MailType1Choice) AddProprietary() *GenericIdentification47 {
	m.Proprietary = new(GenericIdentification47)
	return m.Proprietary
}
