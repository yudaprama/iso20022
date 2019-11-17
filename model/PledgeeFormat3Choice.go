package model

// Choice between formats for the entity to which the financial instruments are pledged.
type PledgeeFormat3Choice struct {

	// Identification of the entity to which the financial instruments are pledged expressed as a code and a BIC.
	TypeAndIdentification *PledgeeTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Identification of the entity to which the financial instruments are pledged expressed as a code and a narrative description.
	Identification *PledgeeTypeAndText1 `xml:"Id"`

	// Identification of the entity to which the financial instruments are pledged expressed as a proprietary type and narrative description.
	Proprietary *GenericIdentification80 `xml:"Prtry"`
}

func (p *PledgeeFormat3Choice) AddTypeAndIdentification() *PledgeeTypeAndAnyBICIdentifier1 {
	p.TypeAndIdentification = new(PledgeeTypeAndAnyBICIdentifier1)
	return p.TypeAndIdentification
}

func (p *PledgeeFormat3Choice) AddIdentification() *PledgeeTypeAndText1 {
	p.Identification = new(PledgeeTypeAndText1)
	return p.Identification
}

func (p *PledgeeFormat3Choice) AddProprietary() *GenericIdentification80 {
	p.Proprietary = new(GenericIdentification80)
	return p.Proprietary
}
