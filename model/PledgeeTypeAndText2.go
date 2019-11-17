package model

// Identification of the entity to which the financial instruments are pledged expressed as a code and a narrative description.
type PledgeeTypeAndText2 struct {

	// Additional information about the entity to which the financial instruments are pledged.
	Identification *RestrictedFINMax30Text `xml:"Id,omitempty"`

	// Entity to which the financial instruments are pledged expressed as a code.
	PledgeeType *PledgeeType1Code `xml:"PldgeeTp"`
}

func (p *PledgeeTypeAndText2) SetIdentification(value string) {
	p.Identification = (*RestrictedFINMax30Text)(&value)
}

func (p *PledgeeTypeAndText2) SetPledgeeType(value string) {
	p.PledgeeType = (*PledgeeType1Code)(&value)
}
