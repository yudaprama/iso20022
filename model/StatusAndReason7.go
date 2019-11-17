package model

// Status and reason of an instructed order.
type StatusAndReason7 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status2Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction14 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason7) AddStatusAndReason() *Status2Choice {
	s.StatusAndReason = new(Status2Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason7) AddTransaction() *Transaction14 {
	newValue := new(Transaction14)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
