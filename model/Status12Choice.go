package model

// Choice of formats for the specification of the status.
type Status12Choice struct {

	// Status of the account opening instruction or account modification instruction.
	Status *AccountManagementStatus1Code `xml:"Sts"`

	// Status of the order is rejected.
	Rejected []*RejectedReason3Choice `xml:"Rjctd"`
}

func (s *Status12Choice) SetStatus(value string) {
	s.Status = (*AccountManagementStatus1Code)(&value)
}

func (s *Status12Choice) AddRejected() *RejectedReason3Choice {
	newValue := new(RejectedReason3Choice)
	s.Rejected = append(s.Rejected, newValue)
	return newValue
}
