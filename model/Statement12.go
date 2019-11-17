package model

// General characteristics related to a statement which reports information.
type Statement12 struct {

	// Indicates whether the statement contains missing instructions only or all instructions.
	StatementType *CorporateActionStatementType1Code `xml:"StmtTp"`

	// Indicates whether the statement report on account holdings for corporate action events is for single account/multiple events or multiple accounts/single event.
	ReportingType *CorporateActionStatementReportingType1Code `xml:"RptgTp"`

	// Reference of the statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Date of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency4Choice `xml:"Frqcy"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *UpdateType2Choice `xml:"UpdTp"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Period during which identification deadline has been set.
	NotificationDeadlinePeriod *DateOrDateTimePeriodChoice `xml:"NtfctnDdlnPrd,omitempty"`
}

func (s *Statement12) SetStatementType(value string) {
	s.StatementType = (*CorporateActionStatementType1Code)(&value)
}

func (s *Statement12) SetReportingType(value string) {
	s.ReportingType = (*CorporateActionStatementReportingType1Code)(&value)
}

func (s *Statement12) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement12) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement12) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement12) AddFrequency() *Frequency4Choice {
	s.Frequency = new(Frequency4Choice)
	return s.Frequency
}

func (s *Statement12) AddUpdateType() *UpdateType2Choice {
	s.UpdateType = new(UpdateType2Choice)
	return s.UpdateType
}

func (s *Statement12) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement12) AddNotificationDeadlinePeriod() *DateOrDateTimePeriodChoice {
	s.NotificationDeadlinePeriod = new(DateOrDateTimePeriodChoice)
	return s.NotificationDeadlinePeriod
}
