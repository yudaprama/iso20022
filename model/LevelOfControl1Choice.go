package model

// Choice of formats for the level of control.
type LevelOfControl1Choice struct {

	// Level of control expressed as a code.
	Code *LevelOfControl1Code `xml:"Cd"`

	// Level of control expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (l *LevelOfControl1Choice) SetCode(value string) {
	l.Code = (*LevelOfControl1Code)(&value)
}

func (l *LevelOfControl1Choice) AddProprietary() *GenericIdentification47 {
	l.Proprietary = new(GenericIdentification47)
	return l.Proprietary
}
