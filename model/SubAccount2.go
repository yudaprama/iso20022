package model

// The subaccount of the safekeeping account
type SubAccount2 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *Max35Text `xml:"Id"`
}

func (s *SubAccount2) SetIdentification(value string) {
	s.Identification = (*Max35Text)(&value)
}
