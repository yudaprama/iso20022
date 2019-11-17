package model

// General characteristics related to a statement which reports information for a defined period.
type Statement8 struct {

	// Reference of the statement.
	Reference *Max35Text `xml:"Ref"`

	// Period on which the statement is reporting.
	StatementPeriod *DatePeriodDetails `xml:"StmtPrd"`

	// Creation date of the statement.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Frequency of the statement.
	Frequency *EventFrequency1Code `xml:"Frqcy,omitempty"`

	// Specifies if the statement is complete or only contains changes.
	UpdateType *StatementUpdateTypeCode `xml:"UpdTp"`

	// Indicates whether there is activity reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Sequential number of the statement.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`
}

func (s *Statement8) SetReference(value string) {
	s.Reference = (*Max35Text)(&value)
}

func (s *Statement8) AddStatementPeriod() *DatePeriodDetails {
	s.StatementPeriod = new(DatePeriodDetails)
	return s.StatementPeriod
}

func (s *Statement8) AddCreationDateTime() *DateAndDateTimeChoice {
	s.CreationDateTime = new(DateAndDateTimeChoice)
	return s.CreationDateTime
}

func (s *Statement8) SetFrequency(value string) {
	s.Frequency = (*EventFrequency1Code)(&value)
}

func (s *Statement8) SetUpdateType(value string) {
	s.UpdateType = (*StatementUpdateTypeCode)(&value)
}

func (s *Statement8) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement8) SetReportNumber(value string) {
	s.ReportNumber = (*Max5NumericText)(&value)
}
