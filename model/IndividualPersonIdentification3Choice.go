package model

// Choice of formats for the identification of an individual person.
type IndividualPersonIdentification3Choice struct {

	// Identification of the person that owns the account.
	IdentificationNumber *GenericIdentification81 `xml:"IdNb"`

	// Name of the person that owns the account.
	PersonName *IndividualPerson35 `xml:"PrsnNm"`
}

func (i *IndividualPersonIdentification3Choice) AddIdentificationNumber() *GenericIdentification81 {
	i.IdentificationNumber = new(GenericIdentification81)
	return i.IdentificationNumber
}

func (i *IndividualPersonIdentification3Choice) AddPersonName() *IndividualPerson35 {
	i.PersonName = new(IndividualPerson35)
	return i.PersonName
}
