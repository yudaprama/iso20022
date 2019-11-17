package model

// Status and reason of an instructed order.
type StatusAndReason1 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status2Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction7 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason1) AddStatusAndReason() *Status2Choice {
	s.StatusAndReason = new(Status2Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason1) AddTransaction() *Transaction7 {
	newValue := new(Transaction7)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
