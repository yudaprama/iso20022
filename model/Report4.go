package model

// General characteristics of the report.
type Report4 struct {

	// Sequential number of the report.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Identification of the SecuritiesStatementQuery message sent to request this statement.
	QueryReference *Max35Text `xml:"QryRef,omitempty"`

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId,omitempty"`

	// Date and time of the report.
	ReportDateTime *DateAndDateTimeChoice `xml:"RptDtTm"`

	// Preparation date and time of the report.
	CreationDateTime *DateAndDateTimeChoice `xml:"CreDtTm,omitempty"`

	// Previous report date and time.
	PreviousReportDateTime *DateAndDateTimeChoice `xml:"PrvsRptDtTm,omitempty"`

	// Specifies the frequency of the report.
	Frequency *Frequency8Choice `xml:"Frqcy"`

	// Specifies whether the report is complete or contains changes only.
	UpdateType *UpdateType4Choice `xml:"UpdTp"`

	// Specifies the type of balance on which the report is prepared.
	ReportBasis *StatementBasis6Choice `xml:"RptBsis"`

	// Period for which the report is given.
	ReportPeriod *DatePeriodDetails `xml:"RptPrd,omitempty"`

	// Specifies the source of the report.
	ReportSource *StatementSource1Choice `xml:"RptSrc,omitempty"`

	// Indicates whether the report is audited or not.
	AuditedIndicator *YesNoIndicator `xml:"AudtdInd,omitempty"`

	// Indicates whether there is activity or an information update reported in the report.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd,omitempty"`
}

func (r *Report4) SetReportNumber(value string) {
	r.ReportNumber = (*Max5NumericText)(&value)
}

func (r *Report4) SetQueryReference(value string) {
	r.QueryReference = (*Max35Text)(&value)
}

func (r *Report4) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *Report4) AddReportDateTime() *DateAndDateTimeChoice {
	r.ReportDateTime = new(DateAndDateTimeChoice)
	return r.ReportDateTime
}

func (r *Report4) AddCreationDateTime() *DateAndDateTimeChoice {
	r.CreationDateTime = new(DateAndDateTimeChoice)
	return r.CreationDateTime
}

func (r *Report4) AddPreviousReportDateTime() *DateAndDateTimeChoice {
	r.PreviousReportDateTime = new(DateAndDateTimeChoice)
	return r.PreviousReportDateTime
}

func (r *Report4) AddFrequency() *Frequency8Choice {
	r.Frequency = new(Frequency8Choice)
	return r.Frequency
}

func (r *Report4) AddUpdateType() *UpdateType4Choice {
	r.UpdateType = new(UpdateType4Choice)
	return r.UpdateType
}

func (r *Report4) AddReportBasis() *StatementBasis6Choice {
	r.ReportBasis = new(StatementBasis6Choice)
	return r.ReportBasis
}

func (r *Report4) AddReportPeriod() *DatePeriodDetails {
	r.ReportPeriod = new(DatePeriodDetails)
	return r.ReportPeriod
}

func (r *Report4) AddReportSource() *StatementSource1Choice {
	r.ReportSource = new(StatementSource1Choice)
	return r.ReportSource
}

func (r *Report4) SetAuditedIndicator(value string) {
	r.AuditedIndicator = (*YesNoIndicator)(&value)
}

func (r *Report4) SetActivityIndicator(value string) {
	r.ActivityIndicator = (*YesNoIndicator)(&value)
}
