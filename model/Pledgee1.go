package model

// Identifies the entity to which the financial instruments are pledged.
type Pledgee1 struct {

	// Unique identification of the party.
	PledgeeTypeAndIdentification *PledgeeFormat3Choice `xml:"PldgeeTpAndId,omitempty"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *Pledgee1) AddPledgeeTypeAndIdentification() *PledgeeFormat3Choice {
	p.PledgeeTypeAndIdentification = new(PledgeeFormat3Choice)
	return p.PledgeeTypeAndIdentification
}

func (p *Pledgee1) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
