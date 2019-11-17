package model

// Choice of identification of an individual person.
type IndividualPersonIdentificationChoice struct {

	// Information related to an identification, eg, party identification or account identification.
	IdentificationNumber *GenericIdentification10 `xml:"IdNb"`

	// Human entity, as distinguished from a corporate entity (which is sometimes referred to as an 'artificial person').
	PersonName *IndividualPerson4 `xml:"PrsnNm"`
}

func (i *IndividualPersonIdentificationChoice) AddIdentificationNumber() *GenericIdentification10 {
	i.IdentificationNumber = new(GenericIdentification10)
	return i.IdentificationNumber
}

func (i *IndividualPersonIdentificationChoice) AddPersonName() *IndividualPerson4 {
	i.PersonName = new(IndividualPerson4)
	return i.PersonName
}
