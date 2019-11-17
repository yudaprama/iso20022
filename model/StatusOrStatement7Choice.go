package model

// Specifies the status or statement that is requested.
type StatusOrStatement7Choice struct {

	// Identify the status advice and the transaction for which the status advice was requested.
	StatusAdvice *DocumentNumber12 `xml:"StsAdvc"`

	// Identify the statement/report that was requested.
	Statement *DocumentNumber13 `xml:"Stmt"`
}

func (s *StatusOrStatement7Choice) AddStatusAdvice() *DocumentNumber12 {
	s.StatusAdvice = new(DocumentNumber12)
	return s.StatusAdvice
}

func (s *StatusOrStatement7Choice) AddStatement() *DocumentNumber13 {
	s.Statement = new(DocumentNumber13)
	return s.Statement
}
