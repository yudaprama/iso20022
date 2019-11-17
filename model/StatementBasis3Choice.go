package model

// Choice of format for the statement basis.
type StatementBasis3Choice struct {

	// Statement basis expressed as an ISO 20022 code.
	Code *StatementBasis1Code `xml:"Cd"`

	// Statement basis expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *StatementBasis3Choice) SetCode(value string) {
	s.Code = (*StatementBasis1Code)(&value)
}

func (s *StatementBasis3Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
