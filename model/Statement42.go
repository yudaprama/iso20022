package model

// Characteristics of the statement.
type Statement42 struct {

	// Date or period of the statement.
	StatementDateOrPeriod *DateAndPeriod1Choice `xml:"StmtDtOrPrd,omitempty"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis7Choice `xml:"StmtBsis,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementType *StatementType5Choice `xml:"StmtTp,omitempty"`
}

func (s *Statement42) AddStatementDateOrPeriod() *DateAndPeriod1Choice {
	s.StatementDateOrPeriod = new(DateAndPeriod1Choice)
	return s.StatementDateOrPeriod
}

func (s *Statement42) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement42) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement42) AddStatementBasis() *StatementBasis7Choice {
	s.StatementBasis = new(StatementBasis7Choice)
	return s.StatementBasis
}

func (s *Statement42) AddStatementType() *StatementType5Choice {
	s.StatementType = new(StatementType5Choice)
	return s.StatementType
}
