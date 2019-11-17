package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson12 struct {

	// Name by which a party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which a person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of an individual, for example, non-pernament resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType1 `xml:"CtryAndResdtlSts,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification11 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson12) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson12) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson12) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType1 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType1)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson12) AddOtherIdentification() *GenericIdentification11 {
	newValue := new(GenericIdentification11)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}
