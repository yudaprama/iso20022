package model

// Reason for the status of the transaction.
type StatusReason4Choice struct {

	// Reason for the status in a coded form.
	Code *FinancingStatusReason1Code `xml:"Cd"`

	// Reason for the status not catered for by the available codes.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *StatusReason4Choice) SetCode(value string) {
	s.Code = (*FinancingStatusReason1Code)(&value)
}

func (s *StatusReason4Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
