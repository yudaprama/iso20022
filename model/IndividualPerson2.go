package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson2 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification8 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson2) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson2) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson2) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson2) AddOtherIdentification() *GenericIdentification8 {
	newValue := new(GenericIdentification8)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}
