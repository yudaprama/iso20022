package model

// Choice of formats for the identification of an individual person.
type IndividualPersonIdentification1Choice struct {

	// Identification of the person that owns the account.
	IdentificationNumber *GenericIdentification81 `xml:"IdNb"`

	// Name of the person that owns the account.
	PersonName *IndividualPerson30 `xml:"PrsnNm"`
}

func (i *IndividualPersonIdentification1Choice) AddIdentificationNumber() *GenericIdentification81 {
	i.IdentificationNumber = new(GenericIdentification81)
	return i.IdentificationNumber
}

func (i *IndividualPersonIdentification1Choice) AddPersonName() *IndividualPerson30 {
	i.PersonName = new(IndividualPerson30)
	return i.PersonName
}
