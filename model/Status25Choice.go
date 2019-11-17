package model

// Choice of formats for the specification of the status.
type Status25Choice struct {

	// Status code.
	Status *AccountManagementStatus1Code `xml:"Sts"`

	// Status of the account management instruction is rejected.
	Rejected []*RejectionReason31 `xml:"Rjctd"`
}

func (s *Status25Choice) SetStatus(value string) {
	s.Status = (*AccountManagementStatus1Code)(&value)
}

func (s *Status25Choice) AddRejected() *RejectionReason31 {
	newValue := new(RejectionReason31)
	s.Rejected = append(s.Rejected, newValue)
	return newValue
}
