package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson8 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max35Text `xml:"Nm"`

	// First name of a person.
	GivenName *Max35Text `xml:"GvnNm"`

	// Specifies the terms used to formally address a person.
	NamePrefix *NamePrefix1Code `xml:"NmPrfx,omitempty"`

	// Additional information about a person that follows a person's name, eg, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Specifies the gender of the person.
	Gender *GenderCode `xml:"Gndr,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Number assigned by a social security agency.
	SocialSecurityNumber *Max35Text `xml:"SclSctyNb,omitempty"`

	// Postal address of a party.
	IndividualInvestorAddress *PostalAddress1 `xml:"IndvInvstrAdr"`
}

func (i *IndividualPerson8) SetName(value string) {
	i.Name = (*Max35Text)(&value)
}

func (i *IndividualPerson8) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson8) SetNamePrefix(value string) {
	i.NamePrefix = (*NamePrefix1Code)(&value)
}

func (i *IndividualPerson8) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson8) SetGender(value string) {
	i.Gender = (*GenderCode)(&value)
}

func (i *IndividualPerson8) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson8) SetSocialSecurityNumber(value string) {
	i.SocialSecurityNumber = (*Max35Text)(&value)
}

func (i *IndividualPerson8) AddIndividualInvestorAddress() *PostalAddress1 {
	i.IndividualInvestorAddress = new(PostalAddress1)
	return i.IndividualInvestorAddress
}
