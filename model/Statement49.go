package model

// Characteristics of the statement.
type Statement49 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Period for the statement.
	StatementPeriod *Period2Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement49) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement49) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement49) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement49) AddStatementPeriod() *Period2Choice {
	s.StatementPeriod = new(Period2Choice)
	return s.StatementPeriod
}

func (s *Statement49) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement49) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement49) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}
