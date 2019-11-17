package model

// Choice of format for the statement basis.
type StatementBasis9Choice struct {

	// Statement basis expressed as an ISO 20022 code.
	Code *StatementBasis1Code `xml:"Cd"`

	// Statement basis expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *StatementBasis9Choice) SetCode(value string) {
	s.Code = (*StatementBasis1Code)(&value)
}

func (s *StatementBasis9Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
