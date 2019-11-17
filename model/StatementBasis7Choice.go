package model

// Choice of format for the statement basis.
type StatementBasis7Choice struct {

	// Statement basis expressed as an ISO 20022 code.
	Code *StatementBasis1Code `xml:"Cd"`

	// Statement basis expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *StatementBasis7Choice) SetCode(value string) {
	s.Code = (*StatementBasis1Code)(&value)
}

func (s *StatementBasis7Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
