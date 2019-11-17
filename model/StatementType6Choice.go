package model

// Choice of format for the statement type.
type StatementType6Choice struct {

	// Statement type expressed as an ISO 20022 code.
	Code *SecuritiesStatementType1Code `xml:"Cd"`

	// Statement type expressed as a proprietary code.
	Proprietary *GenericIdentification47 `xml:"Prtry"`
}

func (s *StatementType6Choice) SetCode(value string) {
	s.Code = (*SecuritiesStatementType1Code)(&value)
}

func (s *StatementType6Choice) AddProprietary() *GenericIdentification47 {
	s.Proprietary = new(GenericIdentification47)
	return s.Proprietary
}
