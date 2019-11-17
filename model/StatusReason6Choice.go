package model

// Specifies the reason for the status of the transaction.
type StatusReason6Choice struct {

	// Reason for the status, as published in an external reason code list.
	Code *ExternalStatusReason1Code `xml:"Cd"`

	// Reason for the status, in a proprietary form.
	Proprietary *Max35Text `xml:"Prtry"`
}

func (s *StatusReason6Choice) SetCode(value string) {
	s.Code = (*ExternalStatusReason1Code)(&value)
}

func (s *StatusReason6Choice) SetProprietary(value string) {
	s.Proprietary = (*Max35Text)(&value)
}
