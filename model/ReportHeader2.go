package model

// Specifies generic information about an investigation report.
type ReportHeader2 struct {

	// Point to point reference as assigned by the case assigner to unambiguously identify the case status report.
	Identification *Max35Text `xml:"Id"`

	// Party reporting the status of the investigation case.
	From *Party7Choice `xml:"Fr"`

	// Party to which the status of the case is reported.
	To *Party7Choice `xml:"To"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (r *ReportHeader2) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportHeader2) AddFrom() *Party7Choice {
	r.From = new(Party7Choice)
	return r.From
}

func (r *ReportHeader2) AddTo() *Party7Choice {
	r.To = new(Party7Choice)
	return r.To
}

func (r *ReportHeader2) SetCreationDateTime(value string) {
	r.CreationDateTime = (*ISODateTime)(&value)
}
