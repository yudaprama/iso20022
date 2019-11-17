package model

// Identification of infrastructure which may be a component of a clearing house and which facilitates clearing and settlement for its members by standing between the buyer and the seller. It may net transactions and it substitutes itself as settlement counterparty for each position.
type PlaceOfClearingIdentification1 struct {

	// Unique identification of the place of clearing.
	Identification *AnyBICIdentifier `xml:"Id,omitempty"`

	// Legal entity identification as an alternate identification for a place of clearing.
	LEI *LEIIdentifier `xml:"LEI,omitempty"`
}

func (p *PlaceOfClearingIdentification1) SetIdentification(value string) {
	p.Identification = (*AnyBICIdentifier)(&value)
}

func (p *PlaceOfClearingIdentification1) SetLEI(value string) {
	p.LEI = (*LEIIdentifier)(&value)
}
