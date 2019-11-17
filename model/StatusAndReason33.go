package model

// Status and reason of an instructed order.
type StatusAndReason33 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status23Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction56 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason33) AddStatusAndReason() *Status23Choice {
	s.StatusAndReason = new(Status23Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason33) AddTransaction() *Transaction56 {
	newValue := new(Transaction56)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
