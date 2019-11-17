package model

// Choice of formats for the specification of the status.
type Status14Choice struct {

	// Status of the account opening instruction or account modification instruction.
	Status *AccountManagementStatus1Code `xml:"Sts"`

	// Status of the order is rejected.
	Rejected []*RejectionReason16 `xml:"Rjctd"`
}

func (s *Status14Choice) SetStatus(value string) {
	s.Status = (*AccountManagementStatus1Code)(&value)
}

func (s *Status14Choice) AddRejected() *RejectionReason16 {
	newValue := new(RejectionReason16)
	s.Rejected = append(s.Rejected, newValue)
	return newValue
}
