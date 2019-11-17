package model

// Choice of format for the statement basis.
type StatementBasis8Choice struct {

	// Statement basis expressed as an ISO 20022 code.
	Code *StatementBasis2Code `xml:"Cd"`

	// Statement basis expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *StatementBasis8Choice) SetCode(value string) {
	s.Code = (*StatementBasis2Code)(&value)
}

func (s *StatementBasis8Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
