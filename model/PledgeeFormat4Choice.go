package model

// Choice between formats for the entity to which the financial instruments are pledged.
type PledgeeFormat4Choice struct {

	// Identification of the entity to which the financial instruments are pledged expressed as a code and a BIC.
	TypeAndIdentification *PledgeeTypeAndAnyBICIdentifier1 `xml:"TpAndId"`

	// Identification of the entity to which the financial instruments are pledged expressed as a code and a narrative description.
	Identification *PledgeeTypeAndText2 `xml:"Id"`

	// Identification of the entity to which the financial instruments are pledged expressed as a proprietary type and narrative description.
	Proprietary *GenericIdentification85 `xml:"Prtry"`
}

func (p *PledgeeFormat4Choice) AddTypeAndIdentification() *PledgeeTypeAndAnyBICIdentifier1 {
	p.TypeAndIdentification = new(PledgeeTypeAndAnyBICIdentifier1)
	return p.TypeAndIdentification
}

func (p *PledgeeFormat4Choice) AddIdentification() *PledgeeTypeAndText2 {
	p.Identification = new(PledgeeTypeAndText2)
	return p.Identification
}

func (p *PledgeeFormat4Choice) AddProprietary() *GenericIdentification85 {
	p.Proprietary = new(GenericIdentification85)
	return p.Proprietary
}
