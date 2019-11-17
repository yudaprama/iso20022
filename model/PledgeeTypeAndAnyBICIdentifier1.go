package model

// Identification of the entity to which the financial instruments are pledged expressed as a code and a BIC.
type PledgeeTypeAndAnyBICIdentifier1 struct {

	// Identification of the entity to which the financial instruments are pledged, expressed as a BIC.
	Identification *AnyBICIdentifier `xml:"Id"`

	// Entity to which the financial instruments are pledged expressed as a code.
	PledgeeType *PledgeeType1Code `xml:"PldgeeTp"`
}

func (p *PledgeeTypeAndAnyBICIdentifier1) SetIdentification(value string) {
	p.Identification = (*AnyBICIdentifier)(&value)
}

func (p *PledgeeTypeAndAnyBICIdentifier1) SetPledgeeType(value string) {
	p.PledgeeType = (*PledgeeType1Code)(&value)
}
