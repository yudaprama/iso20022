package model

// Specifies the type of report.
type ReportType1 struct {

	// Specifies whether the pushed through baseline is a new one or an amendment or a resubmission.
	Type *ReportType1Code `xml:"Tp"`
}

func (r *ReportType1) SetType(value string) {
	r.Type = (*ReportType1Code)(&value)
}
