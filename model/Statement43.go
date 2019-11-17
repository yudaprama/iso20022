package model

// Characteristics of the statement.
type Statement43 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency25Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement43) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement43) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement43) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement43) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement43) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement43) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement43) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}
