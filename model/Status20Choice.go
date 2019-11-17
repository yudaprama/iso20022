package model

// Choice of formats for the specification of the status.
type Status20Choice struct {

	// Status of the account opening instruction or account modification instruction.
	Status *AccountManagementStatus1Code `xml:"Sts"`

	// Status of the order is rejected.
	Rejected []*RejectionReason31 `xml:"Rjctd"`
}

func (s *Status20Choice) SetStatus(value string) {
	s.Status = (*AccountManagementStatus1Code)(&value)
}

func (s *Status20Choice) AddRejected() *RejectionReason31 {
	newValue := new(RejectionReason31)
	s.Rejected = append(s.Rejected, newValue)
	return newValue
}
