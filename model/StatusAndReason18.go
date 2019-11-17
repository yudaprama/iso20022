package model

// Status and reason of an instructed order.
type StatusAndReason18 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status15Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction35 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason18) AddStatusAndReason() *Status15Choice {
	s.StatusAndReason = new(Status15Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason18) AddTransaction() *Transaction35 {
	newValue := new(Transaction35)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
