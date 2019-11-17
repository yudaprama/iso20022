package model

// Provides the parameters of the report.
type ReportParameters2 struct {

	// Unique identification of the report.
	ReportIdentification *Max35Text `xml:"RptId"`

	// Date (and time) at which the report was created.
	ReportDateAndTime *DateAndDateTimeChoice `xml:"RptDtAndTm"`

	// Frequency of the report.
	Frequency *EventFrequency6Code `xml:"Frqcy"`

	// Indicates the currency used for the calculation of the guarantee fund.
	ReportCurrency *CurrencyCode `xml:"RptCcy"`

	// Indicates the date of calculation of the deficit (if any).
	CalculationDate *ISODateTime `xml:"ClctnDt,omitempty"`
}

func (r *ReportParameters2) SetReportIdentification(value string) {
	r.ReportIdentification = (*Max35Text)(&value)
}

func (r *ReportParameters2) AddReportDateAndTime() *DateAndDateTimeChoice {
	r.ReportDateAndTime = new(DateAndDateTimeChoice)
	return r.ReportDateAndTime
}

func (r *ReportParameters2) SetFrequency(value string) {
	r.Frequency = (*EventFrequency6Code)(&value)
}

func (r *ReportParameters2) SetReportCurrency(value string) {
	r.ReportCurrency = (*CurrencyCode)(&value)
}

func (r *ReportParameters2) SetCalculationDate(value string) {
	r.CalculationDate = (*ISODateTime)(&value)
}
