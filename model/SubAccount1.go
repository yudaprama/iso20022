package model

// Business relationship between two entities; one entity is the account owner, the other entity is the account servicer.
type SubAccount1 struct {

	// Unique and unambiguous identification for the account between the account owner and the account servicer.
	Identification *AccountIdentification1 `xml:"Id"`

	// Name of the account. It provides an additional means of identification, and is designated by the account servicer in agreement with the account owner.
	Name *Max35Text `xml:"Nm,omitempty"`

	// Specifies additional properties of the account.
	Characteristic *Max35Text `xml:"Chrtc,omitempty"`
}

func (s *SubAccount1) AddIdentification() *AccountIdentification1 {
	s.Identification = new(AccountIdentification1)
	return s.Identification
}

func (s *SubAccount1) SetName(value string) {
	s.Name = (*Max35Text)(&value)
}

func (s *SubAccount1) SetCharacteristic(value string) {
	s.Characteristic = (*Max35Text)(&value)
}
