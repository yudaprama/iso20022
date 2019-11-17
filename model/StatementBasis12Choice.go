package model

// Choice of format for the statement basis.
type StatementBasis12Choice struct {

	// Statement basis expressed as an ISO 20022 code.
	Code *StatementBasis2Code `xml:"Cd"`

	// Statement basis expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *StatementBasis12Choice) SetCode(value string) {
	s.Code = (*StatementBasis2Code)(&value)
}

func (s *StatementBasis12Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
