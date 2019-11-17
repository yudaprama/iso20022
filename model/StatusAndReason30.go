package model

// Status and reason of an instructed order.
type StatusAndReason30 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status23Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction48 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason30) AddStatusAndReason() *Status23Choice {
	s.StatusAndReason = new(Status23Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason30) AddTransaction() *Transaction48 {
	newValue := new(Transaction48)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
