package model

// Status and reason of an instructed order.
type StatusAndReason16 struct {

	// Status and reason for the transaction.
	StatusAndReason *Status9Choice `xml:"StsAndRsn"`

	// Details of the transactions reported.
	Transaction []*Transaction28 `xml:"Tx,omitempty"`
}

func (s *StatusAndReason16) AddStatusAndReason() *Status9Choice {
	s.StatusAndReason = new(Status9Choice)
	return s.StatusAndReason
}

func (s *StatusAndReason16) AddTransaction() *Transaction28 {
	newValue := new(Transaction28)
	s.Transaction = append(s.Transaction, newValue)
	return newValue
}
