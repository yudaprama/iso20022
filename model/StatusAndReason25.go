package model

// Status and reason of an instructed order.
type StatusAndReason25 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status15Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction40 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason25) AddStatusAndReason() *Status15Choice {
	s.StatusAndReason = new(Status15Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason25) AddTransaction() *Transaction40 {
	newValue := new(Transaction40)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
