package model

// Specifies the source used to generate the reporting.
type ReportingSource1Choice struct {

	// Reporting source, as published in an external reporting source code list.
	Code *ExternalReportingSource1Code `xml:"Cd"`

	// Reporting source, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (r *ReportingSource1Choice) SetCode(value string) {
	r.Code = (*ExternalReportingSource1Code)(&value)
}

func (r *ReportingSource1Choice) SetProprietary(value string) {
	r.Proprietary = (*Max35Text)(&value)
}
