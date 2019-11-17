package model

// Provides the parameters of the report.
type ReportParameters3 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date (and time) and time of the report.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Currency used for the calculation of the margin.
	ReportCurrency *CurrencyCode `xml:"RptCcy"`

	// Date of calculation of the margin.
	CalculationDateAndTime *ISODateTime `xml:"ClctnDtAndTm"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Sequential number of the report.
	ReportNumber *Exact5NumericText `xml:"RptNb,omitempty"`
}

func (r *ReportParameters3) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters3) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters3) SetReportCurrency(value string) {
	r.ReportCurrency = (*CurrencyCode)(&value)
}

func (r *ReportParameters3) SetCalculationDateAndTime(value string) {
	r.CalculationDateAndTime = (*ISODateTime)(&value)
}

func (r *ReportParameters3) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters3) SetReportNumber(value string) {
	r.ReportNumber = (*Exact5NumericText)(&value)
}
