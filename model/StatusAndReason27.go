package model

// Status and reason of an instructed order.
type StatusAndReason27 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status18Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction45 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason27) AddStatusAndReason() *Status18Choice {
	s.StatusAndReason = new(Status18Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason27) AddTransaction() *Transaction45 {
	newValue := new(Transaction45)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
