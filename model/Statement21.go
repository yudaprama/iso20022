package model

// Characteristics of the statement.
type Statement21 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis3Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`
}

func (s *Statement21) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement21) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement21) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement21) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement21) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement21) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement21) AddStatementBasis() *StatementBasis3Choice {
	s.StatementBasis = new(StatementBasis3Choice)
	return s.StatementBasis
}

func (s *Statement21) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement21) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}
