package model

// Status and reason of an instructed order.
type StatusAndReason2 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status2Choice `xml:"StsAndRsn"`
}

func (s *StatusAndReason2) AddStatusAndReason() *Status2Choice {
	s.StatusAndReason = new(Status2Choice)
	return s.StatusAndReason
}
