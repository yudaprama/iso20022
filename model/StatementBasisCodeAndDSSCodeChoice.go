package model

// Choice between formats for the statement basis.
type StatementBasisCodeAndDSSCodeChoice struct {

	// Statement basis expressed as a code.
	StatementBasisAsCode *StatementBasis1Code `xml:"StmtBsisAsCd"`

	// Statement basis expressed as a data source scheme and a code used within the data source scheme.
	StatementBasisAsDSS *GenericIdentification7 `xml:"StmtBsisAsDSS"`
}

func (s *StatementBasisCodeAndDSSCodeChoice) SetStatementBasisAsCode(value string) {
	s.StatementBasisAsCode = (*StatementBasis1Code)(&value)
}

func (s *StatementBasisCodeAndDSSCodeChoice) AddStatementBasisAsDSS() *GenericIdentification7 {
	s.StatementBasisAsDSS = new(GenericIdentification7)
	return s.StatementBasisAsDSS
}
