package model

// Characteristics of the statement.
type Statement14 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp,omitempty"`

	// Specifies whether the statement is sorted by status or transaction.
	StatementStructure *StatementStructure1Code `xml:"StmtStr"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement14) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement14) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement14) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement14) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement14) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement14) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement14) SetStatementStructure(value string) {
	s.StatementStructure = (*StatementStructure1Code)(&value)
}

func (s *Statement14) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}
