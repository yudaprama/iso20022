package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson30 struct {

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which the person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`
}

func (i *IndividualPerson30) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson30) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson30) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson30) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson30) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}
