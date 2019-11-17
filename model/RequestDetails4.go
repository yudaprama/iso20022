package model

// Details of one or several keys of the request.
type RequestDetails4 struct {

	// Key for which the specific data is returned, for example, a BIC.
	Key *Max35Text `xml:"Key"`

	// Data being returned.
	ReportData []*ReportParameter1 `xml:"RptData,omitempty"`
}

func (r *RequestDetails4) SetKey(value string) {
	r.Key = (*Max35Text)(&value)
}

func (r *RequestDetails4) AddReportData() *ReportParameter1 {
	newValue := new(ReportParameter1)
	r.ReportData = append(r.ReportData, newValue)
	return newValue
}
