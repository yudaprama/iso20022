package model

// Choice between the identification of a status or statement query.
type StatusOrStatement3Choice struct {

	// Identify the status advice and the transaction for which the status advice was requested.
	StatusAdvice *DocumentNumber6 `xml:"StsAdvc"`

	// Identify the statement/report that was requested.
	Statement *DocumentNumber1 `xml:"Stmt"`
}

func (s *StatusOrStatement3Choice) AddStatusAdvice() *DocumentNumber6 {
	s.StatusAdvice = new(DocumentNumber6)
	return s.StatusAdvice
}

func (s *StatusOrStatement3Choice) AddStatement() *DocumentNumber1 {
	s.Statement = new(DocumentNumber1)
	return s.Statement
}
