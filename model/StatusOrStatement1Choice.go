package model

// Choice between the identification of a status or statement query.
type StatusOrStatement1Choice struct {

	// Identify the status advice and the transaction for which the status advice was requested.
	StatusAdvice *DocumentNumber2 `xml:"StsAdvc"`

	// Identify the statement/report that was requested.
	Statement *DocumentNumber1 `xml:"Stmt"`
}

func (s *StatusOrStatement1Choice) AddStatusAdvice() *DocumentNumber2 {
	s.StatusAdvice = new(DocumentNumber2)
	return s.StatusAdvice
}

func (s *StatusOrStatement1Choice) AddStatement() *DocumentNumber1 {
	s.Statement = new(DocumentNumber1)
	return s.Statement
}
