package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson14 struct {

	// Name received at birth, eg, maiden name.
	BirthName *Max35Text `xml:"BirthNm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Unique and unambiguous identification of a person, eg passport.
	Identification *PersonIdentification2 `xml:"Id,omitempty"`

	// Postal address of a party.
	Address *PostalAddress1 `xml:"Adr,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingParty *PartyIdentification9Choice `xml:"EmplngPty,omitempty"`
}

func (i *IndividualPerson14) SetBirthName(value string) {
	i.BirthName = (*Max35Text)(&value)
}

func (i *IndividualPerson14) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson14) AddIdentification() *PersonIdentification2 {
	i.Identification = new(PersonIdentification2)
	return i.Identification
}

func (i *IndividualPerson14) AddAddress() *PostalAddress1 {
	i.Address = new(PostalAddress1)
	return i.Address
}

func (i *IndividualPerson14) AddEmployingParty() *PartyIdentification9Choice {
	i.EmployingParty = new(PartyIdentification9Choice)
	return i.EmployingParty
}
