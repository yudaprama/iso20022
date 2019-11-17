package model

// Status and reason of an instructed order.
type StatusAndReason9 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status9Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction20 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason9) AddStatusAndReason() *Status9Choice {
	s.StatusAndReason = new(Status9Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason9) AddTransaction() *Transaction20 {
	newValue := new(Transaction20)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
