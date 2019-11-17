package model

// Characteristics of the statement.
type Statement32 struct {

	// Reference common to all pages of a statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Frequency of the statement.
	Frequency *Frequency1Code `xml:"Frqcy"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`
}

func (s *Statement32) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement32) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement32) SetFrequency(value string) {
	s.Frequency = (*Frequency1Code)(&value)
}

func (s *Statement32) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}
