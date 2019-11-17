package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson4 struct {

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Second name of a person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`
}

func (i *IndividualPerson4) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson4) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson4) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson4) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson4) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}
