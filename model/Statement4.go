package model

// General characteristics related to a statement which reports information for a precise date.
type Statement4 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Preparation date of the statement
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *FrequencyCodeAndDSSCodeChoice `xml:"Frqcy,omitempty"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *StatementUpdateTypeCodeAndDSSCodeChoice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasisCodeAndDSSCodeChoice `xml:"StmtBsis,omitempty"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Indicates whether the statement is audited.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`
}

func (s *Statement4) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement4) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement4) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement4) AddFrequency() *FrequencyCodeAndDSSCodeChoice {
	s.Frequency = new(FrequencyCodeAndDSSCodeChoice)
	return s.Frequency
}

func (s *Statement4) AddUpdateType() *StatementUpdateTypeCodeAndDSSCodeChoice {
	s.UpdateType = new(StatementUpdateTypeCodeAndDSSCodeChoice)
	return s.UpdateType
}

func (s *Statement4) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement4) AddStatementBasis() *StatementBasisCodeAndDSSCodeChoice {
	s.StatementBasis = new(StatementBasisCodeAndDSSCodeChoice)
	return s.StatementBasis
}

func (s *Statement4) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement4) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}
