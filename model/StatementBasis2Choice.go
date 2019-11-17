package model

// Choice of format for the statement basis.
type StatementBasis2Choice struct {

	// Statement basis expressed as an ISO 20022 code.
	Code *StatementBasis2Code `xml:"Cd"`

	// Statement basis expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *StatementBasis2Choice) SetCode(value string) {
	s.Code = (*StatementBasis2Code)(&value)
}

func (s *StatementBasis2Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
