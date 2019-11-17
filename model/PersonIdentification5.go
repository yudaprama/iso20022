package model

// Unique and unambiguous way to identify a person.
type PersonIdentification5 struct {

	// Date and place of birth of a person.
	DateAndPlaceOfBirth *DateAndPlaceOfBirth `xml:"DtAndPlcOfBirth,omitempty"`

	// Unique identification of a person, as assigned by an institution, using an identification scheme.
	Other []*GenericPersonIdentification1 `xml:"Othr,omitempty"`
}

func (p *PersonIdentification5) AddDateAndPlaceOfBirth() *DateAndPlaceOfBirth {
	p.DateAndPlaceOfBirth = new(DateAndPlaceOfBirth)
	return p.DateAndPlaceOfBirth
}

func (p *PersonIdentification5) AddOther() *GenericPersonIdentification1 {
	newValue := new(GenericPersonIdentification1)
	p.Other = append(p.Other, newValue)
	return newValue
}
