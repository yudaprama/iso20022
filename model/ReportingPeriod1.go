package model

// Specifies the requested reporting period.
type ReportingPeriod1 struct {

	// Specifies a date range.
	FromToDate *DatePeriodDetails1 `xml:"FrToDt"`

	// Specifies a time range.
	FromToTime *TimePeriodDetails1 `xml:"FrToTm"`

	// Specifies whether all matching items need to be reported or only those items that are new or have changed since the last similar request was made.
	Type *QueryType3Code `xml:"Tp"`
}

func (r *ReportingPeriod1) AddFromToDate() *DatePeriodDetails1 {
	r.FromToDate = new(DatePeriodDetails1)
	return r.FromToDate
}

func (r *ReportingPeriod1) AddFromToTime() *TimePeriodDetails1 {
	r.FromToTime = new(TimePeriodDetails1)
	return r.FromToTime
}

func (r *ReportingPeriod1) SetType(value string) {
	r.Type = (*QueryType3Code)(&value)
}
