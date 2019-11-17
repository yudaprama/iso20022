package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson35 struct {

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm"`

	// Gender of the person.
	Gender *Gender1Code `xml:"Gndr,omitempty"`

	// Date on which the person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`
}

func (i *IndividualPerson35) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson35) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson35) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson35) SetGender(value string) {
	i.Gender = (*Gender1Code)(&value)
}

func (i *IndividualPerson35) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}
