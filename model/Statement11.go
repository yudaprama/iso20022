package model

// Characteristics of the statement.
type Statement11 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis2Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement11) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement11) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement11) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement11) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement11) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement11) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement11) AddStatementBasis() *StatementBasis2Choice {
	s.StatementBasis = new(StatementBasis2Choice)
	return s.StatementBasis
}

func (s *Statement11) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement11) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}
