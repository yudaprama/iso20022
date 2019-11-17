package model

// Provides the parameters of the report.
type ReportParameters4 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date and time of the report .
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`
}

func (r *ReportParameters4) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters4) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}
