package model

// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
type IndividualPerson33 struct {

	// Term used to address the person.
	NamePrefix *NamePrefix1Choice `xml:"NmPrfx,omitempty"`

	// First name of the person.
	GivenName *Max35Text `xml:"GvnNm,omitempty"`

	// Second name of the person.
	MiddleName *Max35Text `xml:"MddlNm,omitempty"`

	// Name by which the party is known and which is usually used to identify that person.
	Name *Max350Text `xml:"Nm"`

	// Additional information about the person that follows a person's name, for example, qualification such as Doctor of Philosophy (PhD).
	NameSuffix *Max35Text `xml:"NmSfx,omitempty"`

	// Gender of the person.
	Gender *Gender1Code `xml:"Gndr,omitempty"`

	// Date on which the person was born.
	BirthDate *ISODate `xml:"BirthDt,omitempty"`

	// Country where the person was born.
	CountryOfBirth *CountryCode `xml:"CtryOfBirth,omitempty"`

	// Province where the person was born.
	ProvinceOfBirth *Max35Text `xml:"PrvcOfBirth,omitempty"`

	// City where the person was born.
	CityOfBirth *Max35Text `xml:"CityOfBirth,omitempty"`

	// Name of the occupation or job of the person.
	Profession *Max35Text `xml:"Prfssn,omitempty"`

	// Information related to an address to be inserted, updated or deleted.
	ModifiedPostalAddress []*ModificationScope34 `xml:"ModfdPstlAdr,omitempty"`

	// Citizenship information to be inserted or deleted.
	ModifiedCitizenship []*ModificationScope39 `xml:"ModfdCtznsh,omitempty"`

	// Organisation represented by a person, or for which a person works.
	EmployingCompany *Max140Text `xml:"EmplngCpny,omitempty"`

	// Title of the function.
	BusinessFunction *Max35Text `xml:"BizFctn,omitempty"`

	// Specifies if due diligence checks on the political exposure of the investor or account servicer have been carried out and whether these checks are national or foreign. (A politically exposed person is someone who has been entrusted with a prominent public function, or an individual who is closely related to such a person.)
	PoliticallyExposedPersonType *PoliticalExposureType1Choice `xml:"PltclyXpsdPrsnTp,omitempty"`

	// Date of death.
	DeathDate *ISODate `xml:"DthDt,omitempty"`

	// Civil status of the individual person.
	CivilStatus *CivilStatus1Choice `xml:"CvlSts,omitempty"`

	// Highest level of education reached by the individual person.
	EducationLevel *Max35Text `xml:"EdctnLvl,omitempty"`

	// Information related to the person.
	FamilyInformation *PersonalInformation1 `xml:"FmlyInf,omitempty"`
}

func (i *IndividualPerson33) AddNamePrefix() *NamePrefix1Choice {
	i.NamePrefix = new(NamePrefix1Choice)
	return i.NamePrefix
}

func (i *IndividualPerson33) SetGivenName(value string) {
	i.GivenName = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetMiddleName(value string) {
	i.MiddleName = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetName(value string) {
	i.Name = (*Max350Text)(&value)
}

func (i *IndividualPerson33) SetNameSuffix(value string) {
	i.NameSuffix = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetGender(value string) {
	i.Gender = (*Gender1Code)(&value)
}

func (i *IndividualPerson33) SetBirthDate(value string) {
	i.BirthDate = (*ISODate)(&value)
}

func (i *IndividualPerson33) SetCountryOfBirth(value string) {
	i.CountryOfBirth = (*CountryCode)(&value)
}

func (i *IndividualPerson33) SetProvinceOfBirth(value string) {
	i.ProvinceOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetCityOfBirth(value string) {
	i.CityOfBirth = (*Max35Text)(&value)
}

func (i *IndividualPerson33) SetProfession(value string) {
	i.Profession = (*Max35Text)(&value)
}

func (i *IndividualPerson33) AddModifiedPostalAddress() *ModificationScope34 {
	newValue := new(ModificationScope34)
	i.ModifiedPostalAddress = append(i.ModifiedPostalAddress, newValue)
	return newValue
}

func (i *IndividualPerson33) AddModifiedCitizenship() *ModificationScope39 {
	newValue := new(ModificationScope39)
	i.ModifiedCitizenship = append(i.ModifiedCitizenship, newValue)
	return newValue
}

func (i *IndividualPerson33) SetEmployingCompany(value string) {
	i.EmployingCompany = (*Max140Text)(&value)
}

func (i *IndividualPerson33) SetBusinessFunction(value string) {
	i.BusinessFunction = (*Max35Text)(&value)
}

func (i *IndividualPerson33) AddPoliticallyExposedPersonType() *PoliticalExposureType1Choice {
	i.PoliticallyExposedPersonType = new(PoliticalExposureType1Choice)
	return i.PoliticallyExposedPersonType
}

func (i *IndividualPerson33) SetDeathDate(value string) {
	i.DeathDate = (*ISODate)(&value)
}

func (i *IndividualPerson33) AddCivilStatus() *CivilStatus1Choice {
	i.CivilStatus = new(CivilStatus1Choice)
	return i.CivilStatus
}

func (i *IndividualPerson33) SetEducationLevel(value string) {
	i.EducationLevel = (*Max35Text)(&value)
}

func (i *IndividualPerson33) AddFamilyInformation() *PersonalInformation1 {
	i.FamilyInformation = new(PersonalInformation1)
	return i.FamilyInformation
}
