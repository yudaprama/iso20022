package model

// Choice of formats for liability.
type Liability1Choice struct {

	// Liability expressed as a code.
	Code *Liability1Code `xml:"Cd"`

	// Liability expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (l *Liability1Choice) SetCode(value string) {
	l.Code = (*Liability1Code)(&value)
}

func (l *Liability1Choice) AddProprietary() *GenericIdentification47 {
	l.Proprietary = new(GenericIdentification47)
	return l.Proprietary
}
