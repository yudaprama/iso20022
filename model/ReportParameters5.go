package model

// Provides the parameters of the report.
type ReportParameters5 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date (and time) at which the report was created.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Indicates the currency used for the calculation of the guarantee fund.
	ReportCurrency *ActiveCurrencyCode `xml:"RptCcy"`

	// Indicates the date of calculation of the deficit (if any).
	CalculationDate *ISODateTime `xml:"ClctnDt,omitempty"`
}

func (r *ReportParameters5) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters5) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters5) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters5) SetReportCurrency(value string) {
	r.ReportCurrency = (*ActiveCurrencyCode)(&value)
}

func (r *ReportParameters5) SetCalculationDate(value string) {
	r.CalculationDate = (*ISODateTime)(&value)
}
