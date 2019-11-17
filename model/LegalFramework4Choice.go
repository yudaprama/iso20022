package model

// Choice of format for the legal framework.
type LegalFramework4Choice struct {

	// Legal framework expressed as an ISO 20022 code.
	Code *LegalFramework1Code `xml:"Cd"`

	// Legal framework expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (l *LegalFramework4Choice) SetCode(value string) {
	l.Code = (*LegalFramework1Code)(&value)
}

func (l *LegalFramework4Choice) AddProprietary() *GenericIdentification47 {
	l.Proprietary = new(GenericIdentification47)
	return l.Proprietary
}
