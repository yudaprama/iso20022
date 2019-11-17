package model

// Choice of format for the statement type.
type StatementType5Choice struct {

	// Statement type expressed as an ISO 20022 code.
	Code *SecuritiesStatementType1Code `xml:"Cd"`

	// Statement type expressed as a proprietary code.
	Proprietary *GenericIdentification30 `xml:"Prtry"`
}

func (s *StatementType5Choice) SetCode(value string) {
	s.Code = (*SecuritiesStatementType1Code)(&value)
}

func (s *StatementType5Choice) AddProprietary() *GenericIdentification30 {
	s.Proprietary = new(GenericIdentification30)
	return s.Proprietary
}
