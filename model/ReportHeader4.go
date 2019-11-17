package model

// Specifies generic information about an investigation report.
type ReportHeader4 struct {

	// Point to point reference as assigned by the case assigner to unambiguously identify the case status report.
	Identification *Max35Text `xml:"Id"`

	// Party reporting the status of the investigation case.
	From *Party12Choice `xml:"Fr"`

	// Party to which the status of the case is reported.
	To *Party12Choice `xml:"To"`

	// Date and time at which the message was created.
	CreationDateTime *ISODateTime `xml:"CreDtTm"`
}

func (r *ReportHeader4) SetIdentification(value string) {
	r.Identification = (*Max35Text)(&value)
}

func (r *ReportHeader4) AddFrom() *Party12Choice {
	r.From = new(Party12Choice)
	return r.From
}

func (r *ReportHeader4) AddTo() *Party12Choice {
	r.To = new(Party12Choice)
	return r.To
}

func (r *ReportHeader4) SetCreationDateTime(value string) {
	r.CreationDateTime = (*ISODateTime)(&value)
}
