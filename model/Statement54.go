package model

// Characteristics of the statement.
type Statement54 struct {

	// Date or period of the statement.
	StatementDateOrPeriod *DateAndPeriod1Choice `xml:"StmtDtOrPrd,omitempty"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis9Choice `xml:"StmtBsis,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementType *StatementType6Choice `xml:"StmtTp,omitempty"`
}

func (s *Statement54) AddStatementDateOrPeriod() *DateAndPeriod1Choice {
	s.StatementDateOrPeriod = new(DateAndPeriod1Choice)
	return s.StatementDateOrPeriod
}

func (s *Statement54) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement54) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement54) AddStatementBasis() *StatementBasis9Choice {
	s.StatementBasis = new(StatementBasis9Choice)
	return s.StatementBasis
}

func (s *Statement54) AddStatementType() *StatementType6Choice {
	s.StatementType = new(StatementType6Choice)
	return s.StatementType
}
