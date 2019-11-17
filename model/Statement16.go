package model

// Characteristics of the statement.
type Statement16 struct {

	// Date or period of the statement.
	StatementDateOrPeriod *DateAndPeriod1Choice `xml:"StmtDtOrPrd,omitempty"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis3Choice `xml:"StmtBsis,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementType *StatementType2Choice `xml:"StmtTp,omitempty"`
}

func (s *Statement16) AddStatementDateOrPeriod() *DateAndPeriod1Choice {
	s.StatementDateOrPeriod = new(DateAndPeriod1Choice)
	return s.StatementDateOrPeriod
}

func (s *Statement16) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement16) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement16) AddStatementBasis() *StatementBasis3Choice {
	s.StatementBasis = new(StatementBasis3Choice)
	return s.StatementBasis
}

func (s *Statement16) AddStatementType() *StatementType2Choice {
	s.StatementType = new(StatementType2Choice)
	return s.StatementType
}
