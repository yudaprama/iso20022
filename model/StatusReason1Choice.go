package model

// Reason for the status of the transaction.
type StatusReason1Choice struct {

	// Reason for the status in a coded form.
	Code *TransactionRejectReason2Code `xml:"Cd"`

	// Reason for the status not catered for by the available codes.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *StatusReason1Choice) SetCode(value string) {
	s.Code = (*TransactionRejectReason2Code)(&value)
}

func (s *StatusReason1Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
