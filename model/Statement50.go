package model

// Characteristics of the statement.
type Statement50 struct {

	// Identification assigned by the portfolio transfer counterpart to unambiguously identify a portfolio transfer notification.
	CounterpartyPortfolioTransferNotificationReference *RestrictedFINXMax16Text `xml:"CtrPtyPrtflTrfNtfctnRef,omitempty"`

	// Sequential number of the report.
	ReportNumber *Number3Choice `xml:"RptNb,omitempty"`

	// Reference common to all pages of a statement.
	StatementIdentification *RestrictedFINXMax16Text `xml:"StmtId,omitempty"`

	// Date and time of the statement.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *UpdateType16Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement50) SetCounterpartyPortfolioTransferNotificationReference(value string) {
	s.CounterpartyPortfolioTransferNotificationReference = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement50) AddReportNumber() *Number3Choice {
	s.ReportNumber = new(Number3Choice)
	return s.ReportNumber
}

func (s *Statement50) SetStatementIdentification(value string) {
	s.StatementIdentification = (*RestrictedFINXMax16Text)(&value)
}

func (s *Statement50) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement50) AddUpdateType() *UpdateType16Choice {
	s.UpdateType = new(UpdateType16Choice)
	return s.UpdateType
}

func (s *Statement50) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}
