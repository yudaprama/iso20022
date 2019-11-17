package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson31 struct {

	// Name by which the party is known and which is usually used to identify that party.
	Name *Max350Text `xml:"Nm,omitempty"`

	// Date on which the person is born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country and residential status of the individual, for example, non-permanent resident.
	CountryAndResidentialStatus *CountryAndResidentialStatusType2 `xml:"CtryAndResdtlSts,omitempty"`

	// Beneficial owner or its designated agent certifies that it complies with any holding or investment restrictions or requirements of the fund.
	BeneficiaryCertificationCompletion *BeneficiaryCertificationCompletion1Code `xml:"BnfcryCertfctnCmpltn,omitempty"`

	// Alternative identification, for example, national registration identification number, passport number, or an account number used to further identify the beneficial owner, for example, a Central Provident Fund (CFP) account as required for Singapore.
	OtherIdentification []*GenericIdentification164 `xml:"OthrId,omitempty"`
}

func (i *IndividualPerson31) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson31) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson31) AddCountryAndResidentialStatus() *CountryAndResidentialStatusType2 {
	i.CountryAndResidentialStatus = new(CountryAndResidentialStatusType2)
	return i.CountryAndResidentialStatus
}

func (i *IndividualPerson31) SetBeneficiaryCertificationCompletion(value string) {
	i.BeneficiaryCertificationCompletion = (*BeneficiaryCertificationCompletion1Code)(&value)
}

func (i *IndividualPerson31) AddOtherIdentification() *GenericIdentification164 {
	newValue := new(GenericIdentification164)
	i.OtherIdentification = append(i.OtherIdentification, newValue)
	return newValue
}
