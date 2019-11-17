package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson29 struct {

	// Term used to address the person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that person.
	Name *Max350Text `xml:"Nm"`

	// Address of the person.
	PostalAddress []*PostalAddress21 `xml:"PstlAdr"`
}

func (i *IndividualPerson29) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson29) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson29) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson29) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson29) AddPostalAddress() *PostalAddress21 {
	newValue := new(PostalAddress21)
	i.PostalAddress = append(i.PostalAddress, newValue)
	return newValue
}
