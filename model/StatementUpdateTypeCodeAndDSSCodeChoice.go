package model

// Choice between formats for the update type.
type StatementUpdateTypeCodeAndDSSCodeChoice struct {

	// Update type expressed as a code.
	StatementUpdateTypeAsCode *StatementUpdateTypeCode `xml:"StmtUpdTpAsCd"`

	// Update type expressed as a data source scheme and a code used within the data source scheme.
	StatementUpdateTypeAsDSS *GenericIdentification7 `xml:"StmtUpdTpAsDSS"`
}

func (s *StatementUpdateTypeCodeAndDSSCodeChoice) SetStatementUpdateTypeAsCode(value string) {
	s.StatementUpdateTypeAsCode = (*StatementUpdateTypeCode)(&value)
}

func (s *StatementUpdateTypeCodeAndDSSCodeChoice) AddStatementUpdateTypeAsDSS() *GenericIdentification7 {
	s.StatementUpdateTypeAsDSS = new(GenericIdentification7)
	return s.StatementUpdateTypeAsDSS
}
