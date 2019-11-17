package model

// Provides acceptance status of the holding item.
type ReportItemStatus1Choice struct {

	// Statement is accepted.
	Accepted *NoReasonCode `xml:"Accptd"`

	// Statement is accepted  with an exception/s.
	AcceptedWithException []*ReportItemStatus1 `xml:"AccptdWthXcptn"`

	// Statement is rejected.
	Rejected *ReportItemStatus1 `xml:"Rjctd"`
}

func (r *ReportItemStatus1Choice) SetAccepted(value string) {
	r.Accepted = (*NoReasonCode)(&value)
}

func (r *ReportItemStatus1Choice) AddAcceptedWithException() *ReportItemStatus1 {
	newValue := new(ReportItemStatus1)
	r.AcceptedWithException = append(r.AcceptedWithException, newValue)
	return newValue
}

func (r *ReportItemStatus1Choice) AddRejected() *ReportItemStatus1 {
	r.Rejected = new(ReportItemStatus1)
	return r.Rejected
}
