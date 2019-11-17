package model

// Choice of formats for the identification of an individual person.
type IndividualPersonIdentification2Choice struct {

	// Identification of a party, such as a tax or social security identifier.
	IdentificationNumber *GenericIdentification81 `xml:"IdNb"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	PersonName *IndividualPerson30 `xml:"PrsnNm"`
}

func (i *IndividualPersonIdentification2Choice) AddIdentificationNumber() *GenericIdentification81 {
	i.IdentificationNumber = new(GenericIdentification81)
	return i.IdentificationNumber
}

func (i *IndividualPersonIdentification2Choice) AddPersonName() *IndividualPerson30 {
	i.PersonName = new(IndividualPerson30)
	return i.PersonName
}
