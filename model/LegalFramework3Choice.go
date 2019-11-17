package model

// Choice of format for the legal framework.
type LegalFramework3Choice struct {

	// Legal framework expressed as an ISO 20022 code.
	Code *LegalFramework1Code `xml:"Cd"`

	// Legal framework expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (l *LegalFramework3Choice) SetCode(value string) {
	l.Code = (*LegalFramework1Code)(&value)
}

func (l *LegalFramework3Choice) AddProprietary() *GenericIdentification30 {
	l.Proprietary = new(GenericIdentification30)
	return l.Proprietary
}
