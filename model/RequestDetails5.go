package model

// Report of the requested data.
type RequestDetails5 struct {

	// Type of data requested, for example, a sub-member BIC.
	Type *Max35Text `xml:"Tp"`

	// Reference to the request for which the report is sent.
	RequestReference *Max35Text `xml:"ReqRef"`

	// Report key and returned data.
	ReportKey []*RequestDetails4 `xml:"RptKey"`
}

func (r *RequestDetails5) SetType(value string) {
	r.Type = (*Max35Text)(&value)
}

func (r *RequestDetails5) SetRequestReference(value string) {
	r.RequestReference = (*Max35Text)(&value)
}

func (r *RequestDetails5) AddReportKey() *RequestDetails4 {
	newValue := new(RequestDetails4)
	r.ReportKey = append(r.ReportKey, newValue)
	return newValue
}
