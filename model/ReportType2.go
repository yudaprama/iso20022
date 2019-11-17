package model

// Specifies the type of report.
type ReportType2 struct {

	// Specifies whether the report is precalculated or current.
	Type *ReportType2Code `xml:"Tp"`
}

func (r *ReportType2) SetType(value string) {
	r.Type = (*ReportType2Code)(&value)
}
