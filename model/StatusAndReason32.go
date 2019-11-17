package model

// Status and reason of an instructed order.
type StatusAndReason32 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status18Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction54 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason32) AddStatusAndReason() *Status18Choice {
	s.StatusAndReason = new(Status18Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason32) AddTransaction() *Transaction54 {
	newValue := new(Transaction54)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
