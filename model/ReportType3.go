package model

// Specifies the type of report.
type ReportType3 struct {

	// Specifies whether the report is for a matched or pre-matched data set.
	Type *InstructionType3Code `xml:"Tp"`
}

func (r *ReportType3) SetType(value string) {
	r.Type = (*InstructionType3Code)(&value)
}
