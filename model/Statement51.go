package model

// Characteristics of the statement.
type Statement51 struct {

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *RestrictedFINXMax16Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Frequency of the statement.
	Frequency *Frequency26Choice `xml:"Frqcy"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp"`

	// Type of balance on which the statement is prepared.
	StatementBasis *StatementBasis9Choice `xml:"StmtBsis"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`

	// Indicates whether the statement is audited or not.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd"`

	// Indicates whether the statement reports holdings at subsafekeeping account level.
	SubAccountIndicator *YesNoIndicator `xml:"SubAcctInd"`

	// Indicates whether the statement contains tax lot details.
	TaxLotIndicator *YesNoIndicator `xml:"TaxLotInd,omitempty"`
}

func (s *Statement51) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement51) SetQueryReference(value string) {
	s.QueryReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement51) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement51) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement51) AddFrequency() *Frequency26Choice {
	s.Frequency = new(Frequency26Choice)
	return s.Frequency
}

func (s *Statement51) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement51) AddStatementBasis() *StatementBasis9Choice {
	s.StatementBasis = new(StatementBasis9Choice)
	return s.StatementBasis
}

func (s *Statement51) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement51) SetAuditedIndicator(value string) {
	s.AuditedIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement51) SetSubAccountIndicator(value string) {
	s.SubAccountIndicator = (*YesNoIndicator)(&value)
}

func (s *Statement51) SetTaxLotIndicator(value string) {
	s.TaxLotIndicator = (*YesNoIndicator)(&value)
}
