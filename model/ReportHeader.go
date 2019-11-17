package model

// Specifies generic information about an investigation report.
type ReportHeader struct {

	// Identification of the report.
	Identification *Max35Text `xml:"Id"`

	// Party reporting the status of the case.
	From *AnyBICIdentifier `xml:"Fr"`

	// Party to which the status of the case is reported.
	To *AnyBICIdentifier `xml:"To"`

	// Creation date and time of the report generation.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (r *ReportHeader) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportHeader) SetFrom(value string) {
	r.From = (*AnyBICIdentifier)(&value)
}

func (r *ReportHeader) SetTo(value string) {
	r.To = (*AnyBICIdentifier)(&value)
}

func (r *ReportHeader) SetCreationDateTime(value string) {
	r.CreationDateTime = (*ISODateTime)(&value)
}
