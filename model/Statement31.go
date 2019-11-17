package model

// Provides statement details such as the account owner identification (ie, the clearing member identification) and optionaly the non clearing member identification, the clearing account or the list of trade legs.
type Statement31 struct {

	// Identification that is common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Date of the statement.
	StatementDateAndTime *DateAndDateTimeChoice `xml:"StmtDtAndTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *StatementUpdateType1Code `xml:"UpdTp"`

	// Frequency of the statement.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Sequential number of the statement.
	ReportNumber *Exact5NumericText `xml:"RptNb,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement31) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement31) AddStatementDateAndTime() *DateAndDateTimeChoice {
	s.StatementDateAndTime = new(DateAndDateTimeChoice)
	return s.StatementDateAndTime
}

func (s *Statement31) SetUpdateType(value string) {
	s.UpdateType = (*StatementUpdateType1Code)(&value)
}

func (s *Statement31) SetFrequency(value string) {
	s.Frequency = (*EventFrequency6Code)(&value)
}

func (s *Statement31) SetReportNumber(value string) {
	s.ReportNumber = (*Exact5NumericText)(&value)
}

func (s *Statement31) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}
