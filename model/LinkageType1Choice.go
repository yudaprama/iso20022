package model

// Choice of format for the linkage type.
type LinkageType1Choice struct {

	// Linkage type expressed as an ISO 20022 code.
	Code *LinkageType1Code `xml:"Cd"`

	// Linkage type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (l *LinkageType1Choice) SetCode(value string) {
	l.Code = (*LinkageType1Code)(&value)
}

func (l *LinkageType1Choice) AddProprietary() *GenericIdentification20 {
	l.Proprietary = new(GenericIdentification20)
	return l.Proprietary
}
