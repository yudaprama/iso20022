package model

// Identifies the entity to which the financial instruments are pledged.
type Pledgee2 struct {

	// Unique identification of the party.
	PledgeeTypeAndIdentification *PledgeeFormat4Choice `xml:"PldgeeTpAndId,omitempty"`

	// Legal entity identification as an alternate identification for a party.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *Pledgee2) AddPledgeeTypeAndIdentification() *PledgeeFormat4Choice {
	p.PledgeeTypeAndIdentification = new(PledgeeFormat4Choice)
	return p.PledgeeTypeAndIdentification
}

func (p *Pledgee2) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
