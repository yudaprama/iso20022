package model

// General characteristics related to a statement which reports information.
type Statement47 struct {

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
	Frequency *Frequency25Choice `xml:"Frqcy"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *UpdateType15Choice `xml:"UpdTp"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Period during which identification deadline has been set.
	NotificationDeadlinePeriod *DateOrDateTimePeriodChoice `xml:"NtfctnDdlnPrd,omitempty"`
}

func (s *Statement47) SetStatementType(value string) {
	s.StatementType = (*CorporateActionStatementType1Code)(&value)
}

func (s *Statement47) SetReportingType(value string) {
	s.ReportingType = (*CorporateActionStatementReportingType1Code)(&value)
}

func (s *Statement47) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement47) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}

func (s *Statement47) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement47) AddFrequency() *Frequency25Choice {
	s.Frequency = new(Frequency25Choice)
	return s.Frequency
}

func (s *Statement47) AddUpdateType() *UpdateType15Choice {
	s.UpdateType = new(UpdateType15Choice)
	return s.UpdateType
}

func (s *Statement47) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement47) AddNotificationDeadlinePeriod() *DateOrDateTimePeriodChoice {
	s.NotificationDeadlinePeriod = new(DateOrDateTimePeriodChoice)
	return s.NotificationDeadlinePeriod
}
