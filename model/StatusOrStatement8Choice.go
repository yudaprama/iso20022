package model

// Specifies the status or statement that is requested.
type StatusOrStatement8Choice struct {

	// Identify the status advice and the transaction for which the status advice was requested.
	StatusAdvice *DocumentNumber15 `xml:"StsAdvc"`

	// Identify the statement/report that was requested.
	Statement *DocumentNumber14 `xml:"Stmt"`
}

func (s *StatusOrStatement8Choice) AddStatusAdvice() *DocumentNumber15 {
	s.StatusAdvice = new(DocumentNumber15)
	return s.StatusAdvice
}

func (s *StatusOrStatement8Choice) AddStatement() *DocumentNumber14 {
	s.Statement = new(DocumentNumber14)
	return s.Statement
}
