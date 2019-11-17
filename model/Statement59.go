package model

// General characteristics related to a statement which reports information for a precise date.
type Statement59 struct {

	// Specifies the business role of the message sender and, therefore, the business relationship between the sender and the receiver (or the interests represented by them, in those cases where another entity is acting on behalf of the sender or receiver). The message is exchanged between two entities, one being the account servicer and the other the account owner, and the message can be used with either one as the sender.
	SenderBusinessRole *SenderBusinessRole1Code `xml:"SndrBizRole"`

	// Sequential number of the report.
	StatementNumber *Number3Choice `xml:"StmtNb,omitempty"`

	// Identification of the query message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Reference common to all pages of the statement.
	StatementIdentification *Max35Text `xml:"StmtId"`

	// Date and time when the statement was created.
	StatementDateTime *DateAndDateTimeChoice `xml:"StmtDtTm"`

	// Date period for which the statement was created.
	StatementPeriod *DatePeriod1Choice `xml:"StmtPrd"`

	// Frequency of the statement.
	Frequency *Frequency22Choice `xml:"Frqcy,omitempty"`

	// Granularity of the frequency used for the reporting.
	FrequencyGranularity *FrequencyGranularityType1Code `xml:"FrqcyGrnlrty,omitempty"`

	// Specifies whether the statement is complete or contains changes only.
	UpdateType *UpdateType4Choice `xml:"UpdTp,omitempty"`

	// Indicates whether there is activity or updated information reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (s *Statement59) SetSenderBusinessRole(value string) {
	s.SenderBusinessRole = (*SenderBusinessRole1Code)(&value)
}

func (s *Statement59) AddStatementNumber() *Number3Choice {
	s.StatementNumber = new(Number3Choice)
	return s.StatementNumber
}

func (s *Statement59) SetQueryReference(value string) {
	s.QueryReference = (*Max35Text)(&value)
}

func (s *Statement59) SetStatementIdentification(value string) {
	s.StatementIdentification = (*Max35Text)(&value)
}

func (s *Statement59) AddStatementDateTime() *DateAndDateTimeChoice {
	s.StatementDateTime = new(DateAndDateTimeChoice)
	return s.StatementDateTime
}

func (s *Statement59) AddStatementPeriod() *DatePeriod1Choice {
	s.StatementPeriod = new(DatePeriod1Choice)
	return s.StatementPeriod
}

func (s *Statement59) AddFrequency() *Frequency22Choice {
	s.Frequency = new(Frequency22Choice)
	return s.Frequency
}

func (s *Statement59) SetFrequencyGranularity(value string) {
	s.FrequencyGranularity = (*FrequencyGranularityType1Code)(&value)
}

func (s *Statement59) AddUpdateType() *UpdateType4Choice {
	s.UpdateType = new(UpdateType4Choice)
	return s.UpdateType
}

func (s *Statement59) SetActivityIndicator(value string) {
	s.ActivityIndicator = (*YesNoIndicator)(&value)
}
