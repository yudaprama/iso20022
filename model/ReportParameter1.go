package model

// Name and value of a parameter being returned.
type ReportParameter1 struct {

	// Name or type of the parameter being returned.
	Name *Max70Text `xml:"Nm"`

	// Value of the parameter being returned.
	Value *Max350Text `xml:"Val"`
}

func (r *ReportParameter1) SetName(value string) {
	r.Name = (*Max70Text)(&value)
}

func (r *ReportParameter1) SetValue(value string) {
	r.Value = (*Max350Text)(&value)
}
