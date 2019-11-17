package model

// Choice of format for the statement type.
type StatementType2Choice struct {

	// Statement type expressed as an ISO 20022 code.
	Code *SecuritiesStatementType1Code `xml:"Cd"`

	// Statement type expressed as a proprietary code.
	Proprietary *GenericIdentification20 `xml:"Prtry"`
}

func (s *StatementType2Choice) SetCode(value string) {
	s.Code = (*SecuritiesStatementType1Code)(&value)
}

func (s *StatementType2Choice) AddProprietary() *GenericIdentification20 {
	s.Proprietary = new(GenericIdentification20)
	return s.Proprietary
}
