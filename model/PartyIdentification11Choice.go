package model

// Choice of identification of a party. The party can be identified by providing a BIC or a MIC or a proprietary code.
type PartyIdentification11Choice struct {

	// Code allocated to a financial or non-financial institution by the ISO 9362 Registration Authority, as described in ISO 9362 "Banking - Banking telecommunication messages - Business identifier code (BIC)".
	BICOrBEI *AnyBICIdentifier `xml:"BICOrBEI"`

	// An infrastructure that is very often a component of a clearinghouse, and facilitates clearing and settlement for its members by standing between the buyer and the seller of a trade. It may net transactions, and substitutes itself as a settlement counterparty to each position.
	CentralCounterparty *MICIdentifier `xml:"CntrlCtrPty"`

	// Unique and unambiguous identifier, as assigned to a financial institution using a proprietary identification scheme.
	ProprietaryIdentification *GenericIdentification1 `xml:"PrtryId"`
}

func (p *PartyIdentification11Choice) SetBICOrBEI(value string) {
	p.BICOrBEI = (*AnyBICIdentifier)(&value)
}

func (p *PartyIdentification11Choice) SetCentralCounterparty(value string) {
	p.CentralCounterparty = (*MICIdentifier)(&value)
}

func (p *PartyIdentification11Choice) AddProprietaryIdentification() *GenericIdentification1 {
	p.ProprietaryIdentification = new(GenericIdentification1)
	return p.ProprietaryIdentification
}
