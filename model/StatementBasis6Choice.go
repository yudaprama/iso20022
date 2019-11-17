package model

// Choice of format for the statement basis.
type StatementBasis6Choice struct {

	// Statement basis expressed in coded form.
	Code *StatementBasis1Code `xml:"Cd"`

	// Statement basis expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *StatementBasis6Choice) SetCode(value string) {
	s.Code = (*StatementBasis1Code)(&value)
}

func (s *StatementBasis6Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
