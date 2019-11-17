package model

// General characteristics related to a statement which reports information for a precise date.
type Report3 struct {

	// Sequential number of the report.
	ReportNumber *Max5NumericText `xml:"RptNb,omitempty"`

	// Gives the name and the reference of the query.
	QueryReference *QueryReference `xml:"QryRef,omitempty"`

	// Reference of the report.
	ReportIdentification *Max35Text `xml:"RptId,omitempty"`

	// Date of the statement.
	ReportDateTime *DateAndDateTime1Choice `xml:"RptDtTm"`

	// Specifies the regularity of an event.
	Frequency *Frequency4Choice `xml:"Frqcy,omitempty"`

	// Indicates whether the report is complete or contains changes only.
	UpdateType *StatementUpdateTypeCodeAndDSSCodeChoice `xml:"UpdTp,omitempty"`

	// Notifies the type of report transmitted.
	NoticeType *GenericIdentification38 `xml:"NtceTp,omitempty"`
}

func (r *Report3) SetReportNumber(value string) {
	r.ReportNumber = (*Max5NumericText)(&value)
}

func (r *Report3) AddQueryReference() *QueryReference {
	r.QueryReference = new(QueryReference)
	return r.QueryReference
}

func (r *Report3) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *Report3) AddReportDateTime() *DateAndDateTime1Choice {
	r.ReportDateTime = new(DateAndDateTime1Choice)
	return r.ReportDateTime
}

func (r *Report3) AddFrequency() *Frequency4Choice {
	r.Frequency = new(Frequency4Choice)
	return r.Frequency
}

func (r *Report3) AddUpdateType() *StatementUpdateTypeCodeAndDSSCodeChoice {
	r.UpdateType = new(StatementUpdateTypeCodeAndDSSCodeChoice)
	return r.UpdateType
}

func (r *Report3) AddNoticeType() *GenericIdentification38 {
	r.NoticeType = new(GenericIdentification38)
	return r.NoticeType
}
