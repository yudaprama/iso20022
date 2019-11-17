package model

// Choice between the identification of a status or statement query.
type StatusOrStatement5Choice struct {

	// Identify the status advice and the transaction for which the status advice was requested.
	StatusAdvice *DocumentNumber9 `xml:"StsAdvc"`

	// Identify the statement/report that was requested.
	Statement *DocumentNumber1 `xml:"Stmt"`
}

func (s *StatusOrStatement5Choice) AddStatusAdvice() *DocumentNumber9 {
	s.StatusAdvice = new(DocumentNumber9)
	return s.StatusAdvice
}

func (s *StatusOrStatement5Choice) AddStatement() *DocumentNumber1 {
	s.Statement = new(DocumentNumber1)
	return s.Statement
}
