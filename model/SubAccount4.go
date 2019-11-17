package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccount4 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification26 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Specifies additional properties of the account.
	Characteristic *Max35Text `xml:"Chrtc,omitempty"`
}

func (s *SubAccount4) AddIdentification() *AccountIdentification26 {
	s.Identification = new(AccountIdentification26)
	return s.Identification
}

func (s *SubAccount4) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SubAccount4) SetCharacteristic(value string) {
	s.Characteristic = (*Max35Text)(&value)
}
