package model

// Information related to the identification of a person.
type PersonalInformation1 struct {

	// Name of the father of the individual person.
	NameOfFather *Max35Text `xml:"NmOfFthr,omitempty"`

	// Maiden (unmarried) name of the mother of the individual person.
	MaidenNameOfMother *Max35Text `xml:"MdnNmOfMthr,omitempty"`

	// Name of the partner of the individual person.
	NameOfPartner *Max35Text `xml:"NmOfPrtnr,omitempty"`
}

func (p *PersonalInformation1) SetNameOfFather(value string) {
	p.NameOfFather = (*Max35Text)(&value)
}

func (p *PersonalInformation1) SetMaidenNameOfMother(value string) {
	p.MaidenNameOfMother = (*Max35Text)(&value)
}

func (p *PersonalInformation1) SetNameOfPartner(value string) {
	p.NameOfPartner = (*Max35Text)(&value)
}
