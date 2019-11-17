package model

// Choice of format for the legal framework.
type LegalFramework1Choice struct {

	// Legal framework expressed as an ISO 20022 code.
	Code *LegalFramework1Code `xml:"Cd"`

	// Legal framework expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (l *LegalFramework1Choice) SetCode(value string) {
	l.Code = (*LegalFramework1Code)(&value)
}

func (l *LegalFramework1Choice) AddProprietary() *GenericIdentification20 {
	l.Proprietary = new(GenericIdentification20)
	return l.Proprietary
}
