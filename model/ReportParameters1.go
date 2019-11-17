package model

// Parameters related to the net position.
type ReportParameters1 struct {

	// After netting, reference that is common to a net transaction to settle and all its underlying trades.
	NetPositionIdentification *Max35Text `xml:"NetPosId"`

	// Date and time of the net position report.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Indicates whether the statement is complete or contains changes only.
	UpdateType *StatementUpdateType1Code `xml:"UpdTp"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Sequential number of the report.
	ReportNumber *Exact5NumericText `xml:"RptNb,omitempty"`

	// Indicates whether there is activity or information update reported in the statement.
	ActivityIndicator *YesNoIndicator `xml:"ActvtyInd"`
}

func (r *ReportParameters1) SetNetPositionIdentification(value string) {
	r.NetPositionIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters1) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters1) SetUpdateType(value string) {
	r.UpdateType = (*StatementUpdateType1Code)(&value)
}

func (r *ReportParameters1) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters1) SetReportNumber(value string) {
	r.ReportNumber = (*Exact5NumericText)(&value)
}

func (r *ReportParameters1) SetActivityIndicator(value string) {
	r.ActivityIndicator = (*YesNoIndicator)(&value)
}
