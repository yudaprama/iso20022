package model

// Choice of format for the linkage type.
type LinkageType4Choice struct {

	// Linkage type expressed as an ISO 20022 code.
	Code *LinkageType1Code `xml:"Cd"`

	// Linkage type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (l *LinkageType4Choice) SetCode(value string) {
	l.Code = (*LinkageType1Code)(&value)
}

func (l *LinkageType4Choice) AddProprietary() *GenericIdentification47 {
	l.Proprietary = new(GenericIdentification47)
	return l.Proprietary
}
