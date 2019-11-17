package model

// Choice of format for the statement souce.
type StatementSource1Choice struct {

	// Report source expressed in coded form.
	Code *StatementSource1Code `xml:"Cd"`

	// Report source expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *StatementSource1Choice) SetCode(value string) {
	s.Code = (*StatementSource1Code)(&value)
}

func (s *StatementSource1Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
